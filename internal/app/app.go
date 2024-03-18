package app

import (
	"github.com/gin-gonic/gin"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/relic"
	"mevway/internal/app/web"
	"mevway/internal/ports"
)

type Application struct {
	NewRelic       *relic.NewRelic
	Engine         *gin.Engine
	Clients        Clients
	Routers        Routers
	WebServer      WebServer
	EventPublisher *mevent.Publisher
}

type Clients struct {
	AccessService services.AccessServiceClient
	GameService   services.MeviusGameServiceClient
	SocialService services.MeviusSocialServiceClient
}

type Routers struct {
	Core          *ports.APIRouter
	PublicRouter  *ports.PublicAPIRouter
	PrivateRouter *ports.PrivateAPIRouter
}

type WebServer struct {
	Server          *web.Server
	UpdatePublisher *web.ServerUpdatePublisher
	UpdateConsumer  *web.ServerUpdateConsumer
}
