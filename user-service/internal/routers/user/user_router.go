package user

import (
	"user-service/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.GET("/", userController.GetUsers)
	}
}
