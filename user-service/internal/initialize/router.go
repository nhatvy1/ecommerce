package initialize

import (
	"net/http"
	"user-service/global"
	"user-service/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
