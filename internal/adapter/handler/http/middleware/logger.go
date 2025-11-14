package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

type LoggingMiddleware struct {
	logger *slog.Logger
}

func NewLoggingMiddleware(logger *slog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (m *LoggingMiddleware) Handle(ctx *gin.Context) {

	req := ctx.Request
	start := time.Now()

	entry := m.logger.With(
		slog.Group("request_attr",
			slog.String("uri", req.RequestURI),
			slog.String("method", req.Method),
			slog.String("addr", req.RemoteAddr),
			slog.String("client.ip", IPFromContext(ctx).String()),
		),
	)

	entry.InfoContext(ctx, "request received")

	ctx.Next()

	if user, err := UserIDFromContext(ctx); err == nil {
		entry = entry.With(slog.String("user.id", user.String()))
	}

	if player, err := PlayerIDFromContext(ctx); err == nil {
		entry = entry.With(slog.String("player.id", player.String()))
	}

	entry = entry.With(
		slog.Int("status", ctx.Writer.Status()),
		slog.Duration("latency", time.Since(start)),
	)

	if len(ctx.Errors) == 0 {
		entry.InfoContext(ctx, "request success")
	} else {
		entry.With("error", ctx.Errors.Last().Error()).
			ErrorContext(ctx, "request failed")
	}

}
