package middleware

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/warehouse"
	"time"
)

func Logger(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		info := warehouse.Log.NewLogInfo()
		info.RequestURI = ctx.Request().RequestURI
		info.StatusCode = 200
		info.Method = ctx.Request().Method

		if err := next(ctx); err != nil {
			errLog := warehouse.Log.NewLogError()
			errLog.Message = err.Error()
			errLog.StatusCode = 500
			errLog.Print()

			return err
		}

		duration := time.Since(ctx.RequestTime()).Milliseconds()
		info.Duration = fmt.Sprintf("%vms", duration)
		info.Print()

		return nil
	}
}