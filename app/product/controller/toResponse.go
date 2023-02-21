package controller

func ToBaseResponse(statusCode int, success bool, message string) *BaseResponse {
	baseResponse := new(BaseResponse)
	baseResponse.StatusCode = statusCode
	baseResponse.Success = success
	baseResponse.Message = message
	return baseResponse
}
