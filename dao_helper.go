package core

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ormEngine *gorm.DB

func GetOrmEngine() *gorm.DB {
	if nil == ormEngine {
		panic(errors.New("Please call method 'code.InitMysql(...)'"))
	}
	return ormEngine
}

func InitSqlite(logLevel logger.LogLevel, file string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	if nil != db.Error {
		panic(db.Error)
	}
	ormEngine = db
}

func InitMysql(logLevel logger.LogLevel, username string, password string, host string, port int, database string, args map[string]string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, database)
	for k, v := range args {
		dsn = dsn + "&" + k + "=" + v
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	if nil != db.Error {
		panic(db.Error)
	}
	ormEngine = db
}
