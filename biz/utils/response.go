package utils

import (
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RespErr(ctx context.Context, c *app.RequestContext, err error) {
	msg := map[string]interface{}{
		"message": "fail",
		"code":    ierror.ErrUnknown,
	}
	if err != nil {
		msg = map[string]interface{}{
			"message": err.Error(),
			"code":    ierror.GetErrorCode(err),
		}
		ilog.EventError(ctx, err, "resp_error", "err_trace", ierror.Trace(err))
	}
	c.JSON(consts.StatusOK, msg)
}

func RespOK(ctx context.Context, c *app.RequestContext, data interface{}, message ...string) {
	msgText := "success"
	if len(message) > 0 {
		msgText = message[0]
	}
	msg := map[string]interface{}{
		"message": msgText,
		"code":    ierror.Success,
	}
	if data != nil {
		msg["data"] = data
	}
	//ilog.EventInfo(ctx, "resp_ok", "data", data)
	c.JSON(200, msg)
}
