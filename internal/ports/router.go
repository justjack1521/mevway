package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevrelic"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"mevway/internal/app/handler"
)

type APIRouterDecoration func(router *gin.Engine)

type APIRouter struct {
	Logger             *logrus.Logger
	NewRelic           *mevrelic.NewRelic
	ServerStatusHandle handler.ServerStatusHandler
	TokenAuthHandle    handler.TokenAuthoriseHandler
	UserRoleHandler    handler.UserRoleHandler
}

type BaseAPIRouter interface {
	client(ctx *gin.Context) string
	device(ctx *gin.Context) string
	HandleTokenAuthorise(ctx *gin.Context)
}

func (a *APIRouter) HandleServerStatus(ctx *gin.Context) {
	a.ServerStatusHandle.Handle(ctx, handler.ServerStatus{})
}

func (a *APIRouter) HandlerAlphaTesterAuthorise(ctx *gin.Context) {
	a.UserRoleHandler.Handle(ctx, handler.UserRole{
		UserID:   a.client(ctx),
		RoleName: "alpha_tester",
	})
}

func (a *APIRouter) ApplyRouterDecorations(router *gin.Engine) {
	router.Use(a.CORSMiddleware)
	//router.Use(a.NewRelicMiddleware)
	//router.Use(a.LoggerMiddleware)
	//router.Use(a.ErrorLogMiddleware)
	router.GET("/status", a.HandleServerStatus)
}

func (a *APIRouter) NewRelicMiddleware(c *gin.Context) {
	txn := a.NewRelic.Application.StartTransaction(c.Request.RequestURI)
	c.Request = c.Request.Clone(newrelic.NewContext(c.Request.Context(), txn))
	defer txn.End()
	txn.SetWebRequestHTTP(c.Request)
	writer := txn.SetWebResponse(c.Writer)
	c.Next()
	for _, err := range c.Errors {
		txn.NoticeError(err)
	}
	writer.WriteHeader(c.Writer.Status())
}

func (a *APIRouter) LoggerMiddleware(c *gin.Context) {

	entry := a.Logger.WithContext(c.Request.Context()).WithFields(logrus.Fields{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"addr":   c.Request.RemoteAddr,
	})
	entry.Info("Executing request")

	c.Next()

	if len(c.Errors) == 0 {
		entry.Info("Request executed")
	} else {
		entry.WithError(c.Errors.Last()).Error("Request failed")
	}

}

func (a *APIRouter) ErrorLogMiddleware(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		a.Logger.WithContext(c.Request.Context()).WithError(err.Err)
	}
}

func (a *APIRouter) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("", "")
}

func (a *APIRouter) HandleTokenAuthorise(ctx *gin.Context) {
	a.TokenAuthHandle.Handle(ctx, handler.TokenAuthorise{
		UserID:   a.client(ctx),
		Bearer:   ctx.GetHeader("Authorization"),
		DeviceID: a.device(ctx),
	})
}

func (a *APIRouter) client(ctx *gin.Context) string {
	return ctx.GetHeader("X-API-CLIENT")
}

func (a *APIRouter) device(ctx *gin.Context) string {
	return ctx.GetHeader("X-API-DEVICE")
}
