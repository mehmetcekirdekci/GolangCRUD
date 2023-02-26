package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/services"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/types"
	"net/http"
)

type resource struct {
	baseService services.BaseService
}

func NewController(baseService services.BaseService) *resource {
	return &resource{
		baseService: baseService,
	}
}

// CreateUser godoc
// @Summary Creates user
// @Tags user
// @Param request body CreateUserRequest true "CreateUserRequest"
// @Success 201 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/user [post]
func (receiver *resource) CreateUser(c echo.Context) error {
	request := new(CreateUserRequest)
	response := new(BaseResponse)
	var resultMessage string
	err := c.Bind(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	err = c.Validate(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	createUserDto := request.FromCreateUserRequestToCreateUserDto()
	resultMessage, err = receiver.baseService.Create(*createUserDto)
	if err != nil {
		response = ToBaseResponse(http.StatusInternalServerError, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}
	response = ToBaseResponse(http.StatusCreated, true, resultMessage)
	return c.JSON(response.StatusCode, response)
}

// UpdateUser godoc
// @Summary Updates user by CustomerId
// @Tags user
// @Param request body UpdateUserRequest true "UpdateUserRequest"
// @Failure 201 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/user [put]
func (receiver *resource) UpdateUser(c echo.Context) error {
	request := new(UpdateUserRequest)
	response := new(BaseResponse)
	var resultMessage string
	err := c.Bind(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	_, err = uuid.Parse(request.CustomerId)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	err = c.Validate(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	updateUserDto := request.FromUpdateUserRequestToUpdateUserDto()
	var updatedUser *types.User
	updatedUser, resultMessage, err = receiver.baseService.Update(*updateUserDto)
	if err != nil && updatedUser == nil || &updatedUser.Id == nil {
		response = ToBaseResponse(http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	if err != nil {
		response = ToBaseResponse(http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.StatusCode, response)
	}

	response = ToBaseResponse(http.StatusCreated, true, resultMessage)
	return c.JSON(response.StatusCode, response)
}

// GetUser godoc
// @Summary Gets user by customerId
// @Tags user
// @Param query string customerId true "customerId"
// @Success 201 {object} GetUserResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/user [get]
func (receiver *resource) GetUser(c echo.Context) error {
	response := new(GetUserResponse)
	customerId := c.QueryParam("customerId")

	_, err := uuid.Parse(customerId)
	if err != nil {
		response = ToGetUserResponse(nil, http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	user, resultMessage, err := receiver.baseService.GetById(customerId)
	if err != nil && user == nil || user.CustomerId == "" {
		response := ToGetUserResponse(user, http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	if err != nil {
		response := ToGetUserResponse(nil, http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	response = ToGetUserResponse(user, http.StatusOK, true, resultMessage)
	return c.JSON(http.StatusOK, response)
}

// DeleteUser godoc
// @Summary Deletes user
// @Tags user
// @Param query string customerId true "customerId"
// @Success 200 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/user [delete]
func (receiver *resource) DeleteUser(c echo.Context) error {
	response := new(BaseResponse)
	customerId := c.QueryParam("customerId")

	_, err := uuid.Parse(customerId)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	user, resultMessage, err := receiver.baseService.Delete(customerId)
	if err != nil && user == nil || user.CustomerId == "" {
		response = ToBaseResponse(http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	if err != nil {
		response := ToBaseResponse(http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.StatusCode, response)
	}

	response = ToBaseResponse(http.StatusOK, true, resultMessage)
	return c.JSON(http.StatusOK, response)
}
