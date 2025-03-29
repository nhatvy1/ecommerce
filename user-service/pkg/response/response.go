package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func SuccessResponse(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    httpCode,
		Message: "Success",
		Data:    data,
	})
}

func ErrResponse(ctx *gin.Context, httpCode int, message interface{}) {
	ctx.JSON(httpCode, ErrorResponseData{
		Code:    httpCode,
		Message: message,
	})
}
