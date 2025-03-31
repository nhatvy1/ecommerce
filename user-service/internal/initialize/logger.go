package initialize

import (
	"user-service/global"
	"user-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
