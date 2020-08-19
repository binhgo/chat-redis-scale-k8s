package util

import "gitlab.ghn.vn/common-projects/go-sdk/sdk"

func NewError(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Error, Message: mess}
}

func NewErrorWithData(mess string, data interface{}) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Error, Message: mess, Data: data}
}

func NewInvalid(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Invalid, Message: mess}
}

func NewForbidden(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Forbidden, Message: mess}
}

func NewInvalidWithData(mess string, data interface{}) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Invalid, Message: mess, Data: data}
}

func NewOk(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Ok, Message: mess}
}

func NewOkWithData(mess string, data interface{}) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Ok, Message: mess, Data: data}
}

func NewNotFound(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.NotFound, Message: mess}
}

func NewUnAuthorized(mess string) *sdk.APIResponse {
	return &sdk.APIResponse{Status: sdk.APIStatus.Unauthorized, Message: mess}
}
