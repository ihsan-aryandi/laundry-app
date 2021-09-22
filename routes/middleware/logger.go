package middleware

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/warehouse"
	"time"
)

func LogRequest(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		info := warehouse.Log.NewLogInfo()
		info.RequestURI = ctx.Request().RequestURI
		info.StatusCode = 200
		info.Method = ctx.Request().Method

		if err := next(ctx); err != nil {
			errLog := warehouse.Log.NewLogError()
			errLog.Message = "Request error"
			errLog.CausedBy = err.Error()
			errLog.RequestURI = ctx.Request().RequestURI
			errLog.StatusCode = 500
			errLog.Method = ctx.Request().Method
			errLog.Print()

			return err
		}

		duration := time.Since(ctx.RequestTime()).Milliseconds()
		info.Duration = fmt.Sprintf("%vms", duration)
		info.Message = "OK"
		info.Print()

		return nil
	}
}