package controller

import (
	"net/http"
	"user-service/internal/services"
	"user-service/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(us services.IUserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	data := uc.userService.GetUserById()

	// if err != nil {
	// 	response.ErrResponse(ctx, http.StatusNotFound, "Not found", nil)
	// }
	response.SuccessResponse(ctx, http.StatusOK, data)
}
