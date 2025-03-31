package user

import (
	"fmt"
	"net/http"
	"user-service/internal/model"
	"user-service/internal/services"
	"user-service/pkg/response"
	"user-service/pkg/utils/validations"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Login = new(UserLogin)

type UserLogin struct{}

func (c *UserLogin) Login(ctx *gin.Context) {
	response.SuccessResponse(ctx, 200, "login")
}

func (c *UserLogin) Register(ctx *gin.Context) {
	body := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, response.Msg[response.ErrCodeParamInvalid])
		return
	}

	if err := validations.ValidateFunc(body); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	data, code, err := services.UserLogin().Register(ctx, &body)
	if err != nil {
		fmt.Println(zap.Error(err))
		response.ErrResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, data)
}

func (c *UserLogin) Verify(ctx *gin.Context) {
	response.SuccessResponse(ctx, 200, "verify")
}
