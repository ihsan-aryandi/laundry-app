package controller

import (
	"fmt"
	"github.com/bandungrhapsody/rhaprouter"
	"laundry/constant"
	"laundry/controller/validation"
	"laundry/errtype"
	"laundry/service"
	"laundry/warehouse"
	"time"
)

/*
	Services, Repositories, Validations
*/
var (
	/*
		Services
	*/
	registerService = service.NewRegisterSvc()
	loginService = service.NewLoginSvc()
	userInfoService = service.NewUserInfoSvc()

	/*
		Repositories
	*/

	/*
		Validations
	*/
	registerValidation = validation.NewRegisterValidation()
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

/*
	API Responses
*/
type APIResponse struct {
	Info    Info        `json:"info"`
	Results interface{} `json:"results"`
}

type Info struct {
	Success    bool   `json:"success"`
	ErrorCode  string `json:"error_code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

/*
	Success Responses
*/
func successResponse(data interface{}, message string) *APIResponse {
	return &APIResponse{
		Info: Info{
			Success:    true,
			StatusCode: 200,
			Message:    message,
			Timestamp:  time.Now(),
		},
		Results: data,
	}
}

func successRetrieveDataResponse(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessGetDataMessage, dataName)
	return successResponse(data, message)
}

func successCreateDataResponse(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessCreateDataMessage, dataName)
	return successResponse(data, message)
}

func successUpdateDataResponse(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessUpdateDataMessage, dataName)
	return successResponse(data, message)
}

func successDeleteDataResponse(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessDeleteDataMessage, dataName)
	return successResponse(data, message)
}

func successRegistrationResponse(data interface{}) *APIResponse {
	return successResponse(data, constant.SuccessRegistrationMessage)
}

func successLoginResponse(data interface{}) *APIResponse {
	return successResponse(data, constant.SuccessLoginMessage)
}

func successLogoutResponse(data interface{}) *APIResponse {
	return successResponse(data, constant.SuccessLogoutMessage)
}

/*
	Error Responses
*/
func errResponse(err *errtype.Error) *APIResponse {
	return &APIResponse{
		Info:    Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
		Results: err.Contents,
	}
}

func errDataNotFoundResponse(dataName string) *APIResponse {
	err := errtype.DataNotFoundError(dataName)
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
		Results: nil,
	}
}

func errInvalidRequestBodyResponse(e error) *APIResponse {
	err := errtype.InvalidRequestBodyError(e)
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
		Results: nil,
	}
}

func errDataAlreadyExistsResponse(dataName string) *APIResponse {
	err := errtype.DataIsExistError(dataName)
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
	}
}

func usernameIsAlreadyRegisteredResponse() *APIResponse {
	err := errtype.UsernameIsAlreadyRegisteredError()
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
	}
}

func loginErrorResponse(err error) *APIResponse {
	errType := errtype.LoginError(err)
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  errType.Code,
			StatusCode: errType.StatusCode,
			Message:    errType.Message,
			Timestamp:  time.Now(),
		},
	}
}

func forbiddenResponse() *APIResponse {
	err := errtype.ForbiddenError()
	return &APIResponse{
		Info: Info{
			Success:    false,
			ErrorCode:  err.Code,
			StatusCode: err.StatusCode,
			Message:    err.Message,
			Timestamp:  time.Now(),
		},
	}
}