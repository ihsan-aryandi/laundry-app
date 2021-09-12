package errtype

import (
	"fmt"
	"laundry/constant"
)

type Error struct {
	Code       string
	StatusCode int
	Message    string
	Error      error
	Contents   interface{}
}

/*
	Status Code : 400
*/
func InvalidFieldsError(contents interface{}) *Error {
	return &Error{
		Code:       constant.ErrInvalidFieldsCode,
		StatusCode: 400,
		Message:    constant.ErrInvalidFieldsMessage,
		Contents:   contents,
	}
}

func InvalidRequestBodyError(err error) *Error {
	return &Error{
		Code:       constant.ErrInvalidRequestBodyCode,
		StatusCode: 400,
		Message:    constant.ErrInvalidRequestBodyMessage,
		Error:      err,
	}
}

func DataNotFoundError(dataName string) *Error {
	return &Error{
		Code:       constant.ErrDataNotFoundCode,
		StatusCode: 400,
		Message:    fmt.Sprintf(constant.ErrDataNotFoundMessage, dataName),
	}
}

func DataIsExistError(dataName string) *Error {
	return &Error{
		Code:       constant.ErrDataIsExistCode,
		StatusCode: 400,
		Message:    fmt.Sprintf(constant.ErrDataIsExistMessage, dataName),
	}
}

func UsernameIsAlreadyRegisteredError() *Error {
	return &Error{
		Code:       constant.ErrUsernameIsAlreadyRegisteredCode,
		StatusCode: 400,
		Message:    constant.ErrUsernameIsAlreadyRegisteredMessage,
	}
}

func LoginError(err error) *Error {
	return &Error{
		Code:       constant.ErrLoginCode,
		StatusCode: 400,
		Message:    constant.ErrLoginMessage,
		Error:      err,
	}
}

func ForbiddenError() *Error {
	return &Error{
		Code:       constant.ErrForbiddenCode,
		StatusCode: 403,
		Message:    constant.ErrForbiddenMessage,
	}
}

func UnauthorizedError(err error) *Error {
	return &Error{
		Code:       constant.ErrUnauthorizedCode,
		StatusCode: 401,
		Message:    constant.ErrUnauthorizedMessage,
		Error:      err,
	}
}

/*
	Status Code : 500
*/
func InternalServerError(err error) *Error {
	return &Error{
		Code:       constant.ErrInternalServerCode,
		StatusCode: 500,
		Message:    constant.ErrInternalServerMessage,
		Error:      err,
	}
}
