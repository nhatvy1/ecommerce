package user

import (
	"user-service/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.GET("/:id", userController.GetUser)
		userRouterPublic.POST("/register", userController.Register)
	}
}
