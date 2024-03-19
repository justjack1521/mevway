package app

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/newrelic/go-agent/v3/newrelic"
	"mevway/internal/app/web"
	"mevway/internal/ports"
)

type Application struct {
	NewRelic       *newrelic.Application
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
