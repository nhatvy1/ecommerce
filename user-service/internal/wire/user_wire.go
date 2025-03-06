//go:build wireinject

package wire

import (
	"user-service/internal/controller"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"github.com/google/wire"
)

// func InitUserRouterHandler() (*controller.UserController, error) {
// 	iUserRepository := repositories.NewUserRepository()
// 	iUserService := services.NewUserService(iUserRepository)
// 	userController := controller.NewUserController(iUserService)

// 	return userController, nil
// }

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repositories.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
