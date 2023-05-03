package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initializes the database instance
func Setup() {
	var err error
	// sql 配置
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin-gorm-mysql?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建DB对象并连接数据库, dbName: 数据库名称, dsn: 数据库连接字
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("models setup err: ", err)
	}

	// TODO: 初始化数据表
	DB.AutoMigrate(&Word{})

	// 连接池
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}
