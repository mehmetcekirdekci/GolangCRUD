package services

import (
	"errors"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/repositories"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type (
	BaseService interface {
		Create(createUserDto types.CreateUserDto) (string, error)
		Update(updateUserDto types.UpdateUserDto) (*types.User, string, error)
		GetById(id string) (*types.User, string, error)
		Delete(customerId string) (*types.User, string, error)
	}

	baseService struct {
		userRepository repositories.UserRepository
	}
)

func NewBaseService(userRepository repositories.UserRepository) BaseService {
	if userRepository == nil {
		return nil
	}
	return &baseService{
		userRepository: userRepository,
	}
}

func (receiver *baseService) Create(createUserDto types.CreateUserDto) (string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	user := createUserDto.FromCreateUserDtoToUser()
	err := create(receiver, *user)
	if err == nil {
		isOperationSuccesful = true
	}
	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return resultMessage, err
}

func (receiver *baseService) Update(updateUserDto types.UpdateUserDto) (*types.User, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	userToBeUpdated, err := getById(receiver, updateUserDto.CustomerId)

	UserNotFound := err == nil && userToBeUpdated == nil
	UserFoundedSuccesfuly := err == nil && userToBeUpdated != nil && userToBeUpdated.CustomerId != ""

	if UserNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if UserFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	if !isOperationSuccesful {
		resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
		return nil, resultMessage, err
	}

	updatedUser := updateUserDto.FromUpdateUserDtoToUser(userToBeUpdated.Id)
	err = update(receiver, *updatedUser)
	if err != nil {
		isOperationSuccesful = false
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return updatedUser, resultMessage, err
}

func (receiver *baseService) GetById(id string) (*types.User, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	user, err := getById(receiver, id)

	UserNotFound := err == nil && user == nil
	UserFoundedSuccesfuly := err == nil && user != nil && user.CustomerId != ""

	if UserNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if UserFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return user, resultMessage, err
}

func create(receiver *baseService, user types.User) error {
	ctx := context.TODO()
	result, err := receiver.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	if result.InsertedID == nil {
		return errors.New(types.ObjectCannotBeInserted)
	}
	return nil
}

func (receiver *baseService) Delete(customerId string) (*types.User, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	userToBeDeleted, err := getById(receiver, customerId)

	UserNotFound := err == nil && userToBeDeleted == nil
	UserFoundedSuccesfuly := err == nil && userToBeDeleted != nil && userToBeDeleted.CustomerId != ""

	if UserNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if UserFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	if !isOperationSuccesful {
		resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
		return nil, resultMessage, err
	}

	err = delete(receiver, userToBeDeleted.Id)
	if err != nil {
		isOperationSuccesful = false
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return userToBeDeleted, resultMessage, err
}

func getById(receiver *baseService, customerId string) (*types.User, error) {
	user, err := receiver.userRepository.GetByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func update(receiver *baseService, user types.User) error {
	result, err := receiver.userRepository.Update(user)
	if err != nil {
		return err
	}
	if result.ModifiedCount <= 0 {
		return errors.New(types.ObjectCannotBeUpdated)
	}
	return nil
}

func delete(receiver *baseService, id primitive.ObjectID) error {
	result, err := receiver.userRepository.Delete(id)
	if err != nil {
		return err
	}
	if result.DeletedCount <= 0 {
		return errors.New(types.ObjectCannotBeDeleted)
	}
	return nil
}
