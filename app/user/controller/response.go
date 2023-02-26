package controller

import "github.com/mehmetcekirdekci/GolangCRUD/app/user/types"

type (
	BaseResponse struct {
		StatusCode int
		Success    bool
		Message    string
	}

	GetUserResponse struct {
		BaseResponse BaseResponse
		Data         *GetUserResponseDataItem
	}
	GetUserResponseDataItem struct {
		User types.User
	}
)
