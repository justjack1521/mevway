package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type RelicMiddleware struct {
	relic *newrelic.Application
}

func NewRelicMiddleware(relic *newrelic.Application) *RelicMiddleware {
	return &RelicMiddleware{relic: relic}
}

func (m *RelicMiddleware) Handle(ctx *gin.Context) {
	var txn = m.relic.StartTransaction(ctx.Request.RequestURI)
	ctx.Request = ctx.Request.Clone(newrelic.NewContext(ctx.Request.Context(), txn))
	defer txn.End()
	txn.SetWebRequestHTTP(ctx.Request)
	writer := txn.SetWebResponse(ctx.Writer)
	ctx.Next()
	for _, err := range ctx.Errors {
		txn.NoticeError(err)
	}
	writer.WriteHeader(ctx.Writer.Status())
}
