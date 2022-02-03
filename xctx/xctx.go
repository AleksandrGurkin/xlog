package xctx

import (
	"context"

	"github.com/AleksandrGurkin/xlog"
	"github.com/AleksandrGurkin/xlog/xempty"
)

var ContextValue = "xlogger"

func SaveLoggerToCtx(ctx context.Context, logger xlog.Logger) context.Context {
	return context.WithValue(ctx, ContextValue, logger)
}

func XLogFromCtx(ctx context.Context) xlog.Logger {
	if log, ok := ctx.Value(ContextValue).(xlog.Logger); ok {
		return log
	}
	return xempty.Empty{}
}
