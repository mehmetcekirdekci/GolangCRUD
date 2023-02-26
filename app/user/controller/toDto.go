package controller

import (
	"github.com/google/uuid"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/types"
)

func (request CreateUserRequest) FromCreateUserRequestToCreateUserDto() *types.CreateUserDto {
	createUserDto := new(types.CreateUserDto)
	createUserDto.CustomerId = uuid.New()
	createUserDto.Name = request.Name
	createUserDto.LastName = request.LastName
	createUserDto.Age = request.Age
	return createUserDto
}

func (request UpdateUserRequest) FromUpdateUserRequestToUpdateUserDto() *types.UpdateUserDto {
	updateUserDto := new(types.UpdateUserDto)
	updateUserDto.CustomerId = request.CustomerId
	updateUserDto.Name = request.Name
	updateUserDto.LastName = request.LastName
	updateUserDto.Age = request.Age
	return updateUserDto
}
