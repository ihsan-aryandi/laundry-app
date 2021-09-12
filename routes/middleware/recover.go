package middleware

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/warehouse"
)

func Recover(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		defer func() {
			if r := recover(); r != nil {
				errLog := warehouse.Log.NewLogError()
				errLog.Message = fmt.Sprintf("%v", r)
				errLog.RequestURI = ctx.Request().RequestURI
				errLog.StatusCode = 500
				errLog.Print()

				ctx.StatusCode(500)
				_ = ctx.JSON(rhaprouter.Map{
					"message": "Something is wrong with server",
				})
			}
		}()

		return next(ctx)
	}
}