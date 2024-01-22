package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ? 全局变量
var Db *gorm.DB // 在 main 函数外定义数据库连接

func ConnectMysql() {

	//? 连接数据库
	var err error
	dsn := "root:199517@tcp(127.0.0.1:3306)/hero-ranking?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("mysql数据库连接成功~")
	if err != nil {
		panic("failed to connect database")
	}

	// 连接池
	sqlDB, err := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
}
