package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init 使用 GORM，初始化数据库连接。
func init() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{}) // 连接 SQLite 数据库
	if err != nil {
		fmt.Printf("数据库连接失败：%v", err)
	}
	db.AutoMigrate(&Problem{}, &User{}) // 迁移数据库 Schema
	// 添加预设账户
	for _, u := range predefinedUsers {
		db.Create(&u)
	}
	DB = db
}
