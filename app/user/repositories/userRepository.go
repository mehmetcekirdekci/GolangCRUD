package repositories

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type (
	UserRepository interface {
		Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
		GetByCustomerId(customerId string) (*types.User, error)
		Update(user types.User) (*mongo.UpdateResult, error)
		Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	}

	userRepository struct {
		usersCollection *mongo.Collection
	}
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		usersCollection: database.Collection(Users),
	}
}

func (receiver *userRepository) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := receiver.usersCollection.InsertOne(ctx, document)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (receiver *userRepository) GetByCustomerId(customerId string) (*types.User, error) {
	var user types.User
	ctx := context.TODO()
	filter := bson.D{{"customerId", customerId}}
	//filter := bson.E{"customerId", customerId}
	err := receiver.usersCollection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (receiver *userRepository) Update(user types.User) (*mongo.UpdateResult, error) {
	ctx := context.TODO()
	filter := bson.D{{"_id", user.Id}}
	replacement := types.User{CustomerId: user.CustomerId, Name: user.Name, LastName: user.LastName, Age: user.Age}
	result, err := receiver.usersCollection.ReplaceOne(ctx, filter, replacement)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (receiver *userRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx := context.TODO()
	filter := bson.D{{"_id", id}}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})
	result, err := receiver.usersCollection.DeleteOne(ctx, filter, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
