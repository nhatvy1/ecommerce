package wire

import (
	"user-service/internal/controller"
	"user-service/internal/repositories"
	"user-service/internal/services"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repositories.NewUserRepository()
	iUserService := services.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)

	return userController, nil
}
