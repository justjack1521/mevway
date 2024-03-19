package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevconn"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/app"
	"mevway/internal/app/handler"
	"mevway/internal/app/web"
	"mevway/internal/ports"
)

func NewApplication(ctx context.Context) app.Application {

	application := app.Application{
		EventPublisher: mevent.NewPublisher(),
	}

	nrlc, err := mevconn.CreateNewRelicConfig()
	if err != nil {
		panic(err)
	}

	relic, err := newrelic.NewApplication(
		newrelic.ConfigAppName(nrlc.ApplicationName()),
		newrelic.ConfigLicense(nrlc.LicenseKey()),
	)

	logger := logrus.New()
	nrl := nrlogrus.NewFormatter(relic, &logrus.TextFormatter{})
	logger.SetFormatter(nrl)
	//logger.AddHook(relic)

	mqc, err := mevconn.CreateRabbitMQConfig()
	if err != nil {
		panic(err)
	}
	mq, err := rabbitmq.NewConn(mqc.Source(), rabbitmq.WithConnectionOptionsLogging)

	access, err := DialToAccessClient()
	if err != nil {
		panic(err)
	}

	game, err := DialToGameClient()
	if err != nil {
		panic(err)
	}

	social, err := DialToSocialClient()
	if err != nil {
		panic(err)
	}

	rank, err := DialToRankClient()
	if err != nil {
		panic(err)
	}

	multi, err := DialToLobbyClient()
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	engine.HandleMethodNotAllowed = false

	svr := web.NewServer(logger, relic).WithUpdatePublisher(mq).WithUpdateConsumer(mq)
	svr.RegisterServiceClient(web.GameClientRouteKey, web.NewGameServiceClientRouter(game))
	svr.RegisterServiceClient(web.SocialClientRouteKey, web.NewSocialServiceClientRouter(social))
	svr.RegisterServiceClient(web.RankingClientRouteKey, web.NewRankServiceClientRouter(rank))
	svr.RegisterServiceClient(web.MultiClientRouteKey, web.NewMultiServiceClientRouter(multi))

	core := &ports.APIRouter{
		Logger:             logger,
		NewRelic:           relic,
		ServerStatusHandle: handler.NewServerStatusHandler(),
		TokenAuthHandle:    handler.NewTokenHandler(access),
		UserRoleHandler:    handler.NewUserRoleHandler(access),
	}

	private := &ports.PrivateAPIRouter{
		BaseAPIRouter:   core,
		BanUserHandler:  handler.NewBanUserHandler(access),
		UserRoleHandler: handler.NewUserRoleHandler(access),
	}

	cache := app.NewCustomerIDMemoryCache()

	public := &ports.PublicAPIRouter{
		BaseAPIRouter:      core,
		LoginUserHandle:    handler.NewLoginHandler(access, cache),
		RegisterUserHandle: handler.NewRegisterUserHandler(access),
		WebsocketHandle:    handler.NewWebSocketHandler(svr),
		PlayerSearchHandle: handler.NewPlayerSearchHandler(access, social, cache),
		UserRoleHandler:    handler.NewUserRoleHandler(access),
	}

	core.ApplyRouterDecorations(engine)
	private.ApplyRouterDecorations(engine)
	public.ApplyRouterDecorations(engine)

	application.NewRelic = relic
	application.Engine = engine

	application.WebServer = app.WebServer{
		Server: svr,
	}
	application.Routers = app.Routers{
		Core:          core,
		PublicRouter:  public,
		PrivateRouter: private,
	}

	return application

}
