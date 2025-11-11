package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type LoggingMiddleware struct {
	logger *slog.Logger
}

func NewLoggingMiddleware(logger *slog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (m *LoggingMiddleware) Handle(ctx *gin.Context) {

	var entry = m.logger.With(
		slog.Group("request_attr",
			slog.String("uri", ctx.Request.RequestURI),
			slog.String("method", ctx.Request.Method),
			slog.String("addr", ctx.Request.RemoteAddr),
			slog.String("client.ip", IPFromContext(ctx).String()),
		),
	)

	entry.InfoContext(ctx, "request received")

	ctx.Next()

	user, err := UserIDFromContext(ctx)
	if err == nil {
		entry.With(slog.String("user.id", user.String()))
	}

	player, err := PlayerIDFromContext(ctx)
	if err == nil {
		entry.With(slog.String("player.id", player.String()))
	}

	if len(ctx.Errors) == 0 {
		entry.InfoContext(ctx, "request success")
	} else {
		entry.With("error", ctx.Errors.Last().Error()).ErrorContext(ctx, "request failed")
	}

}
