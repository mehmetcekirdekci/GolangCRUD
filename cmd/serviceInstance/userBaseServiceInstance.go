package serviceInstance

import (
	repositories "github.com/mehmetcekirdekci/GolangCRUD/app/user/repositories"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/services"
)

func GetUserBaseServiceInstance(userRepository repositories.UserRepository) services.BaseService {
	return getUserBaseServiceInstance(userRepository)
}

func getUserBaseServiceInstance(userRepository repositories.UserRepository) services.BaseService {
	return services.NewBaseService(userRepository)
}
