package model

import (
	"fmt"
	"go_blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB
var err error

func InitDb(){
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}
	sqlDB, _ := db.DB()

	// AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
	_ = db.AutoMigrate(&User{},&Category{},&Article{})
	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
