package main

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrelic"
	"github.com/sirupsen/logrus"
	"io"
	"mevway/internal/adapter/broker"
	"mevway/internal/adapter/database"
	"mevway/internal/adapter/external"
	"mevway/internal/adapter/handler/http"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/rpc"
	"mevway/internal/adapter/handler/web"
	"mevway/internal/adapter/keycloak"
	"mevway/internal/adapter/memory"
	"mevway/internal/adapter/translate"
	"mevway/internal/core/application"
	"mevway/internal/infrastructure/instrumentation/relic"
	"mevway/internal/infrastructure/instrumentation/system"
	"os"
)

func main() {

	var logger = logrus.New()
	var publisher = mevent.NewPublisher(mevent.PublisherWithLogger(logger))

	db, err := database.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	rds, err := memory.NewRedisConnection(context.Background())
	if err != nil {
		panic(err)
	}

	cloak, err := mevconn.NewKeyCloakConfig()
	if err != nil {
		panic(err)
	}

	msq, err := broker.NewRabbitMQConnection()
	if err != nil {
		panic(err)
	}

	nrl, err := mevrelic.NewRelicApplication()
	if err != nil {
		panic(err)
	}

	game, err := rpc.DialToGameClient()
	if err != nil {
		panic(err)
	}

	social, err := rpc.DialToSocialClient()
	if err != nil {
		panic(err)
	}

	rank, err := rpc.DialToRankClient()
	if err != nil {
		panic(err)
	}

	multi, err := rpc.DialToLobbyClient()
	if err != nil {
		panic(err)
	}

	var keyCloakClient = gocloak.NewClient(cloak.Hostname())
	var userRepository = keycloak.NewUserClient(keyCloakClient, cloak)
	var tokenRepository = keycloak.NewTokenClient(keyCloakClient, cloak)
	var patchRepository = database.NewPatchRepository(db)
	var clientRepository = memory.NewClientRepository(rds)
	var socialRepository = external.NewSocialPlayerRepository(social)

	var serviceRouter = application.NewServiceRouter()
	serviceRouter.RegisterSubRouter(rpc.GameClientRouteKey, rpc.NewGameServiceClientRouter(game))
	serviceRouter.RegisterSubRouter(rpc.SocialClientRouteKey, rpc.NewSocialServiceClientRouter(social))
	serviceRouter.RegisterSubRouter(rpc.RankingClientRouteKey, rpc.NewRankServiceClientRouter(rank))
	serviceRouter.RegisterSubRouter(rpc.MultiClientRouteKey, rpc.NewMultiServiceClientRouter(multi))

	var server = application.NewSocketServer(publisher)

	var relicInstrumenter = relic.NewRelicInstrumenter(nrl.Application)
	var messageTranslator = translate.NewProtobufSocketMessageTranslator()
	var socketFactory = web.NewClientFactory(serviceRouter, relicInstrumenter, messageTranslator)

	var statusService = system.NewStatusService()
	var authService = application.NewAuthenticationService(tokenRepository, userRepository, publisher)
	var patchService = application.NewPatchService(patchRepository)
	var searchService = application.NewPlayerSearchService(userRepository, socialRepository)

	var loggerMiddleware = middleware.NewLoggingMiddleware(logger)
	var relicMiddleware = middleware.NewRelicMiddleware(nrl.Application)

	var statusHandler = http.NewStatusHandler(statusService)
	var authHandler = http.NewAuthenticationHandler(authService, tokenRepository)
	var patchHandler = http.NewPatchHandler(patchService)
	var socketHandler = http.NewSocketHandler(server, clientRepository, socketFactory)
	var searchHandler = http.NewSearchHandler(searchService)

	var listeners = []io.Closer{
		broker.NewClientNotificationConsumer(msq, server, messageTranslator),
		broker.NewClientEventPublisher(msq, publisher, translate.NewProtobufSocketEventTranslator()),
		broker.NewUserEventPublisher(msq, publisher, translate.NewProtobufUserEventTranslator()),
	}

	broker.NewClientPersistenceConsumer(publisher, clientRepository)

	go server.Run()

	router, err := http.NewRouter(authHandler, statusHandler, patchHandler, socketHandler, searchHandler, loggerMiddleware.Handle, relicMiddleware.Handle)

	if err := router.Serve(":8080"); err != nil {
		for _, listener := range listeners {
			listener.Close()
		}
		os.Exit(1)
	}

}
