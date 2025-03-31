package initialize

import (
	"database/sql"
	"fmt"
	"time"
	"user-service/global"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMysqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "InitMysql initialization error")

	err = db.Ping()
	checkErrorPanic(err, "Failed to connect to MySQL database")

	// global.Logger.Info("Initializing MySQL Successfully sql")
	global.MySQL_SQLC = db

	SetPool()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb := global.MySQL_SQLC
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
