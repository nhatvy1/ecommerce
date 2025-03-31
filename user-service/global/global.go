package global

import (
	"database/sql"
	"user-service/pkg/logger"
	"user-service/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config     setting.Config
	MySQL_Gorm *gorm.DB
	MySQL_SQLC *sql.DB
	Rdb        *redis.Client
	Logger     *logger.LoggerZap
)
