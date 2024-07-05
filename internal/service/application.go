package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevconn"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrelic"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/adapter/cache"
	"mevway/internal/adapter/database"
	"mevway/internal/adapter/external"
	"mevway/internal/adapter/memory"
	"mevway/internal/app"
	"mevway/internal/app/handler"
	"mevway/internal/app/web"
	"mevway/internal/ports"
)

func NewApplication(ctx context.Context) app.Application {

	application := app.Application{
		EventPublisher: mevent.NewPublisher(),
	}

	logger := logrus.New()

	client, err := memory.NewRedisConnection(ctx)
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	relic, err := mevrelic.NewRelicApplication()
	if err != nil {
		panic(err)
	}
	relic.Attach(logger)

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

	svr := web.NewServer(logger, relic.Application).WithUpdatePublisher(mq).WithUpdateConsumer(mq)
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

	var players = cache.NewCustomerPlayerIDRepository(external.NewCustomerPlayerIDRepository(access), memory.NewCustomerPlayerCache(client))
	var patches = database.NewPatchRepository(db)

	public := &ports.PublicAPIRouter{
		BaseAPIRouter:      core,
		LoginUserHandle:    handler.NewLoginHandler(access),
		RegisterUserHandle: handler.NewRegisterUserHandler(access),
		WebsocketHandle:    handler.NewWebSocketHandler(svr),
		PlayerSearchHandle: handler.NewPlayerSearchHandler(social, players),
		UserRoleHandler:    handler.NewUserRoleHandler(access),
		PatchListHandler:   handler.NewPatchListHandler(patches),
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
