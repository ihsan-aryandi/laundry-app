package response

import (
	"fmt"
	"laundry/constant"
	"laundry/errtype"
	"time"
)

type APIResponse struct {
	Info    Info        `json:"info"`
	Results interface{} `json:"results"`
}

type Info struct {
	Success    bool      `json:"success"`
	ErrorCode  string    `json:"error_code,omitempty"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

/*
	Success Responses
*/
func OK(data interface{}, message string) *APIResponse {
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

func SuccessRetrieveData(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessGetDataMessage, dataName)
	return OK(data, message)
}

func SuccessCreateData(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessCreateDataMessage, dataName)
	return OK(data, message)
}

func SuccessUpdateData(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessUpdateDataMessage, dataName)
	return OK(data, message)
}

func SuccessDeleteData(dataName string, data interface{}) *APIResponse {
	message := fmt.Sprintf(constant.SuccessDeleteDataMessage, dataName)
	return OK(data, message)
}

func SuccessRegistration(data interface{}) *APIResponse {
	return OK(data, constant.SuccessRegistrationMessage)
}

func SuccessLogin(data interface{}) *APIResponse {
	return OK(data, constant.SuccessLoginMessage)
}

func SuccessLogout(data interface{}) *APIResponse {
	return OK(data, constant.SuccessLogoutMessage)
}

/*
	Error Responses
*/
func Error(err *errtype.Error) *APIResponse {
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
