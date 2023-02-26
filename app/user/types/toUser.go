package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (createUserDto CreateUserDto) FromCreateUserDtoToUser() *User {
	user := new(User)
	user.CustomerId = createUserDto.CustomerId.String()
	user.Age = createUserDto.Age
	user.Name = createUserDto.Name
	user.LastName = createUserDto.LastName
	return user
}

func (updateUserDto UpdateUserDto) FromUpdateUserDtoToUser(id primitive.ObjectID) *User {
	user := new(User)
	user.Id = id
	user.CustomerId = updateUserDto.CustomerId
	user.Name = updateUserDto.Name
	user.LastName = updateUserDto.LastName
	user.Age = updateUserDto.Age
	return user
}
