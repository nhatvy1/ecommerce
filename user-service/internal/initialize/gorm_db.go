package initialize

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"user-service/global"
	"user-service/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDbWithGorm() {
	var err error
	db, err := connectionDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	global.MySQL_Gorm = db

	err = performMigration()
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}
}

func connectionDb() (*gorm.DB, error) {
	username := global.Config.Mysql.Username
	password := global.Config.Mysql.Password
	dbName := global.Config.Mysql.DbName
	host := global.Config.Mysql.Host
	port := strconv.Itoa(global.Config.Mysql.Port)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
		// Logger: nil,
	})

	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	sqlDatabase.SetMaxIdleConns(10)
	sqlDatabase.SetMaxOpenConns(25)
	sqlDatabase.SetConnMaxLifetime(time.Hour)

	return databaseConnection, nil
}

func performMigration() error {
	models := []interface{}{
		&models.User{},
	}

	for _, model := range models {
		if err := global.MySQL_Gorm.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}
