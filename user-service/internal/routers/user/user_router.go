package user

import (
	"user-service/internal/controller/user"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {

	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", user.Login.Register)
		userRouterPublic.POST("/login", user.Login.Register)
	}
}
