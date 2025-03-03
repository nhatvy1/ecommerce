package initialize

import (
	validations "user-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()

	InitDbWithGorm()

	validations.InitValidator()

	r := InitRouter()

	return r
}
