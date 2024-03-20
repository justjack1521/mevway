package app

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrelic"
	"mevway/internal/app/web"
	"mevway/internal/ports"
)

type Application struct {
	NewRelic       *mevrelic.NewRelic
	Engine         *gin.Engine
	Routers        Routers
	WebServer      WebServer
	EventPublisher *mevent.Publisher
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
