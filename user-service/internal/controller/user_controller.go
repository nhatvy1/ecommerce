package controller

import (
	"net/http"
	"strconv"
	"user-service/internal/services"
	"user-service/internal/vo"
	"user-service/pkg/response"
	"user-service/pkg/utils/validations"

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

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		response.ErrResponse(ctx, http.StatusBadRequest, response.Msg[response.ErrPlsTryAgainLater], nil)
		return
	}

	data, err := uc.userService.GetUserById(id)

	if err != nil {
		response.ErrResponse(ctx, http.StatusNotFound, "Not found", nil)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, data)
}

func (uc *UserController) Register(ctx *gin.Context) {
	user := &vo.UserRegister{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := validations.ValidateFunc(user); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, "Validation failed", err)
		return
	}

	data, err := uc.userService.Register(user)
	if err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, data)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil || id <= 0 {
		response.ErrResponse(ctx, http.StatusBadRequest, response.Msg[response.ErrPlsTryAgainLater], nil)
		return
	}

	user := vo.UserUpdate{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := validations.ValidateFunc(user); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, "Validation failed", err)
		return
	}

	data, err := uc.userService.UpdateUser(id, &user)
	if err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, data)
}
