package helper

import "github.com/soheilsirousi/golang-web-api/src/api/validations"

type BaseResponse struct {
	Result          any                            `json:"result"`
	Success         bool                           `json:"success"`
	ResultCode      int                            `json:"code"`
	ValidationError *[]validations.ValidationError `json:"validation_error"`
	Error           any                            `json:"error"`
}

func NewBaseResponse(result any, success bool, resultCode int) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func NewBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func NewBaseResponseWithAnyError(result any, success bool, resultCode int, err any) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func NewBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseResponse {
	return &BaseResponse{
		Result:          result,
		Success:         success,
		ResultCode:      resultCode,
		ValidationError: validations.GetValidationError(err),
	}
}
