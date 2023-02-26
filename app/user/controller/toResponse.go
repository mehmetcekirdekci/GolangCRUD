package controller

import "github.com/mehmetcekirdekci/GolangCRUD/app/user/types"

func ToBaseResponse(statusCode int, success bool, message string) *BaseResponse {
	baseResponse := new(BaseResponse)
	baseResponse.StatusCode = statusCode
	baseResponse.Success = success
	baseResponse.Message = message
	return baseResponse
}

func ToGetUserResponse(user *types.User, statusCode int, success bool, message string) *GetUserResponse {
	getUserResponse := new(GetUserResponse)
	getUserResponse.BaseResponse.Success = success
	getUserResponse.BaseResponse.Message = message
	getUserResponse.BaseResponse.StatusCode = statusCode
	if user != nil {
		data := new(GetUserResponseDataItem)
		data.User = *user
		getUserResponse.Data = data
	}
	return getUserResponse
}
