package middleware

import (
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/warehouse"
	"time"
)

func Logger(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		info := warehouse.Log.NewLogInfo()
		info.RequestURI = ctx.Request().RequestURI

		if err := next(ctx); err != nil {
			errLog := warehouse.Log.NewLogError()
			errLog.Message = err.Error()
			errLog.StatusCode = 500
			errLog.Print()

			return err
		}

		info.Duration = time.Since(ctx.RequestTime()).Milliseconds()
		info.Print()

		return nil
	}
}