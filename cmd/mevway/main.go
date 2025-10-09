package main

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrelic"
	slogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
	"log/slog"
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
	"mevway/internal/core/application/subscriber"
	"mevway/internal/infrastructure/broker/rmq"
	"mevway/internal/infrastructure/trace/relic"
	"mevway/internal/infrastructure/trace/system"
	"os"
)

func main() {

	var ctx = context.Background()

	var logger = logrus.New()
	var slogger = slog.New(slogrus.Option{Level: slog.LevelDebug, Logger: logger}.NewLogrusHandler())

	var events = mevent.NewPublisher(mevent.PublisherWithLogger(slogger))

	db, err := database.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	rds, err := memory.NewRedisConnection(ctx)
	if err != nil {
		panic(err)
	}

	cloak, err := mevconn.NewKeyCloakConfig()
	if err != nil {
		panic(err)
	}

	mqc, err := rmq.NewRabbitMQConnection()
	if err != nil {
		panic(err)
	}

	nrl, err := mevrelic.NewRelicApplication()
	if err != nil {
		panic(err)
	}
	nrl.Attach(logger)

	game, err := rpc.DialToGameClient()
	if err != nil {
		panic(err)
	}

	admin, err := rpc.DialToAdminClient()
	if err != nil {
		panic(err)
	}

	model, err := rpc.DialToModelClient()
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
	var tokenRepository = keycloak.NewTokenClient(keyCloakClient, cloak, slogger)
	var patchRepository = database.NewPatchRepository(db)
	var contactRepository = database.NewContactRepository(db)
	var clientRepository = memory.NewClientRepository(rds)
	var socialRepository = external.NewSocialPlayerRepository(social)
	var requestRepository = memory.NewRequestMemoryRepository(rds)
	var progressRepository = database.NewProgressRepository(db)
	var rankRepository = external.NewRankingRepository(rank)

	var serviceRouter = application.NewServiceRouter(slogger, requestRepository)
	serviceRouter.RegisterSubRouter(rpc.GameClientRouteKey, rpc.NewGameServiceClientRouter(game))
	serviceRouter.RegisterSubRouter(rpc.SocialClientRouteKey, rpc.NewSocialServiceClientRouter(social))
	serviceRouter.RegisterSubRouter(rpc.RankingClientRouteKey, rpc.NewRankServiceClientRouter(rank))
	serviceRouter.RegisterSubRouter(rpc.MultiClientRouteKey, rpc.NewMultiServiceClientRouter(multi))

	var adminService = external.NewGameAdminService(admin)
	var modelService = external.NewGameValidateService(model)

	var server = application.NewSocketServer(events)
	var tracer = relic.NewRelicTracer(nrl.Application)

	var messageTranslator = translate.NewProtobufSocketMessageTranslator()
	var socketFactory = web.NewClientFactory(server, serviceRouter, tracer, messageTranslator)

	var statusService = system.NewStatusService()
	var authService = application.NewAuthenticationService(events, tokenRepository)
	var userService = application.NewUserService(events, userRepository)
	var patchService = application.NewPatchService(patchRepository)
	var progressService = application.NewProgressService(progressRepository)
	var searchService = application.NewPlayerSearchService(userRepository, socialRepository)
	var rankService = application.NewRankQueryService(rankRepository)

	var loggerMiddleware = middleware.NewLoggingMiddleware(slogger)
	var relicMiddleware = middleware.NewRelicMiddleware(nrl.Application)
	var patchMiddleware = middleware.NewPatchMiddleware()

	var statusHandler = http.NewStatusHandler(statusService)
	var authHandler = http.NewAuthenticationHandler(authService, tokenRepository)
	var userHandler = http.NewUserHandler(userService)
	var patchHandler = http.NewPatchHandler(patchService)
	var progressHandler = http.NewFeatureHandler(progressService)
	var socketHandler = http.NewSocketHandler(server, clientRepository, socketFactory)
	var searchHandler = http.NewSearchHandler(searchService)
	var adminHandler = http.NewAdminHandler(adminService)
	var modelHandler = http.NewModelHandler(modelService)
	var contactHandler = http.NewContactHandler(contactRepository)
	var rankHandler = http.NewRankHandler(rankService)

	subscriber.NewClientPersistenceSubscriber(events, clientRepository)

	var rmqa = rmq.NewApplicationConnection(mqc, slogger, tracer)

	var consumers = []broker.Consumer{
		rmq.NewClientNotificationConsumer(rmqa, server, messageTranslator),
	}

	var publishers = []broker.Publisher{
		rmq.NewSocketClientEventPublisher(rmqa, events, translate.NewProtobufSocketEventTranslator()),
		rmq.NewUserEventPublisher(rmqa, events, translate.NewProtobufUserEventTranslator()),
	}

	server.Start()

	events.Notify(application.NewStartEvent(ctx))

	router, err := http.NewRouter(authHandler, userHandler, statusHandler, patchHandler, progressHandler, socketHandler, searchHandler, adminHandler, modelHandler, contactHandler, rankHandler, loggerMiddleware.Handle, relicMiddleware.Handle, patchMiddleware.Handle)
	if err := router.Serve(":8080"); err != nil {
		events.Notify(application.NewShutdownEvent(ctx))
		for _, consumer := range consumers {
			consumer.Close()
		}
		for _, publisher := range publishers {
			publisher.Close()
		}
		os.Exit(1)
	}

}
