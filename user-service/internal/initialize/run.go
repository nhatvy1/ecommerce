package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()
	InitDbWithGorm()

	r := InitRouter()
	return r
}
