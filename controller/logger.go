package controller

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/errtype"
	"laundry/warehouse"
)

/*
	Logger
*/
type logHandler func(ctx *rhaprouter.Context) (res *APIResponse, errType *errtype.Error)

func logController(ctx *rhaprouter.Context, controller logHandler) error {
	errLog := warehouse.Log.NewLogError()
	errLog.RequestURI = ctx.Request().RequestURI
	errLog.StatusCode = 500
	errLog.Timestamp = ctx.RequestTime()

	res, errType := controller(ctx)
	if errType != nil {
		errLog.StatusCode = errType.StatusCode
		errLog.ErrorCode = errType.Code
		errLog.Message = errType.Message
		if errType.Error != nil {
			errLog.Message = fmt.Sprintf("%v (%v)", errLog.Message, errType.Error)
		}
		errLog.Print()

		return ctx.JSON(errResponse(errType))
	}

	res.Info.Timestamp = ctx.RequestTime()
	return ctx.JSON(res)
}