package repositoryInstance

import (
	repositories "github.com/mehmetcekirdekci/GolangCRUD/app/user/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserRepositoryInstance(mongo *mongo.Database) repositories.UserRepository {
	return getUserRepositoryInstance(mongo)
}

func getUserRepositoryInstance(mongo *mongo.Database) repositories.UserRepository {
	return repositories.NewUserRepository(mongo)
}
