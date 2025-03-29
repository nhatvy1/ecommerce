package initialize

import (
	"user-service/global"
	"user-service/internal/database"
	"user-service/internal/services"
	"user-service/internal/services/impl/user_impl"
)

func InitServiceInterface() {
	queries := database.New(global.MySQL_SQLC)

	// user service interface
	services.InitUserLogin(user_impl.NewUserLoginImpl(queries))
}
