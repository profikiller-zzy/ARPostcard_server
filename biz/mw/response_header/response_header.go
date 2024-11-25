package response_header

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	LogIDKey = "K_LOG_ID"
)

func RespLog() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
		logID, _ := ctx.Value(LogIDKey).(string)
		c.Response.Header.Set("log_id", logID)
	}
}
