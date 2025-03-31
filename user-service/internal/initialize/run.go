package initialize

import (
	"user-service/pkg/utils/validations"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()

	validations.InitValidator()

	r := InitRouter()

	return r
}
