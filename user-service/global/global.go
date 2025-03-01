package global

import (
	"user-service/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config     setting.Config
	MySQL_Gorm *gorm.DB
)
