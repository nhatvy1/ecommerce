package global

import (
	"database/sql"
	"user-service/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config     setting.Config
	MySQL_Gorm *gorm.DB
	MySQL_SQLC *sql.DB
)
