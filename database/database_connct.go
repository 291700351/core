package database

import (
	"fmt"

	"gorm.io/gorm"
)

// OpenDatabase
//
//	@Description: 打开一个新的数据库连接
//	@param dialector 数据库驱动
//	@param opts 配置
//	@return *gorm.DB
//	@return error
func OpenDatabase(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AutoMigrate
//
//	@Description: 数据库迁移
//	@param db
//	@param autoMigrateTable
//	@return error
func AutoMigrate(db *gorm.DB, autoMigrateTable ...interface{}) error {
	return db.AutoMigrate(autoMigrateTable...)
}

func Close(db *gorm.DB) {
	if nil != db {
		s, err := db.DB()
		if nil != err {
			fmt.Println("When shutting down the database, get the sql.DB failed", err)
		} else {
			if err := s.Close(); nil != err {
				fmt.Println("When shutting down the database error", err)
			}
		}
	}
}
