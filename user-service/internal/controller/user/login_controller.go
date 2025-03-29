package user

import (
	"fmt"
	"net/http"
	"user-service/internal/model"
	"user-service/internal/services"
	"user-service/pkg/response"

	"github.com/gin-gonic/gin"
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

	data, code, err := services.UserLogin().Register(ctx, &body)
	fmt.Println(err)
	if err != nil {
		fmt.Println("err: ", err)
		response.ErrResponse(ctx, code, "ds")
		return
	}
	fmt.Println("err2: ", err)

	response.SuccessResponse(ctx, http.StatusOK, data)
}

func (c *UserLogin) Verify(ctx *gin.Context) {
	response.SuccessResponse(ctx, 200, "verify")
}
