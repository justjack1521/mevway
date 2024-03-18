package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/relic"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/sirupsen/logrus"
	"mevway/internal/app"
	"mevway/internal/app/handler"
	"mevway/internal/app/web"
	"mevway/internal/config"
	"mevway/internal/ports"
)

func NewApplication(ctx context.Context, conf config.Application) app.Application {

	application := app.Application{
		EventPublisher: mevent.NewPublisher(),
	}

	newrelic := relic.NewRelicApplication(conf.NewRelic)

	logger := logrus.New()
	nrl := nrlogrus.NewFormatter(newrelic.Application, &logrus.TextFormatter{})
	logger.SetFormatter(nrl)
	logger.AddHook(newrelic)

	mq, err := conf.RabbitMQ.NewRabbitMQConnection()
	if err != nil {
		panic(err)
	}

	access, err := DialToAccessClient(conf)
	if err != nil {
		panic(err)
	}

	game, err := DialToGameClient(conf)
	if err != nil {
		panic(err)
	}

	social, err := DialToSocialClient(conf)
	if err != nil {
		panic(err)
	}

	rank, err := DialToRankClient(conf)
	if err != nil {
		panic(err)
	}

	multi, err := DialToLobbyClient(conf)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	engine.HandleMethodNotAllowed = false

	svr := web.NewServer(logger, newrelic.Application).WithUpdatePublisher(mq).WithUpdateConsumer(mq)
	svr.RegisterServiceClient(web.GameClientRouteKey, web.NewGameServiceClientRouter(game))
	svr.RegisterServiceClient(web.SocialClientRouteKey, web.NewSocialServiceClientRouter(social))
	svr.RegisterServiceClient(web.RankingClientRouteKey, web.NewRankServiceClientRouter(rank))
	svr.RegisterServiceClient(web.MultiClientRouteKey, web.NewMultiServiceClientRouter(multi))

	core := &ports.APIRouter{
		Logger:             logger,
		NewRelic:           newrelic.Application,
		ServerStatusHandle: handler.NewServerStatusHandler(),
		TokenAuthHandle:    handler.NewTokenHandler(access),
		UserRoleHandler:    handler.NewUserRoleHandler(access),
	}

	priv := &ports.PrivateAPIRouter{
		BaseAPIRouter:   core,
		BanUserHandler:  handler.NewBanUserHandler(access),
		UserRoleHandler: handler.NewUserRoleHandler(access),
	}

	cache := app.NewCustomerIDMemoryCache()

	pub := &ports.PublicAPIRouter{
		BaseAPIRouter:      core,
		LoginUserHandle:    handler.NewLoginHandler(access, cache),
		RegisterUserHandle: handler.NewRegisterUserHandler(access),
		WebsocketHandle:    handler.NewWebSocketHandler(svr),
		PlayerSearchHandle: handler.NewPlayerSearchHandler(access, social, cache),
		UserRoleHandler:    handler.NewUserRoleHandler(access),
	}

	core.ApplyRouterDecorations(engine)
	priv.ApplyRouterDecorations(engine)
	pub.ApplyRouterDecorations(engine)

	application.NewRelic = newrelic
	application.Engine = engine

	application.WebServer = app.WebServer{
		Server: svr,
	}
	application.Clients = app.Clients{
		AccessService: access,
		GameService:   game,
		SocialService: social,
	}
	application.Routers = app.Routers{
		Core:          core,
		PublicRouter:  pub,
		PrivateRouter: priv,
	}

	return application

}
