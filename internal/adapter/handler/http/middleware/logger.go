package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	logger *logrus.Logger
}

func NewLoggingMiddleware(logger *logrus.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (m *LoggingMiddleware) Handle(ctx *gin.Context) {

	entry := m.logger.WithContext(ctx.Request.Context()).WithFields(logrus.Fields{
		"uri":    ctx.Request.RequestURI,
		"method": ctx.Request.Method,
		"addr":   ctx.Request.RemoteAddr,
	})
	entry.Info("Request received")

	ctx.Next()

	if len(ctx.Errors) == 0 {
		entry.Info("Request success")
	} else {
		entry.WithError(ctx.Errors.Last()).Error("Request failure")
	}

}
