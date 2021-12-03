package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"<-:create"` // 允许读和创建
	Age  int    `gorm:"-"`         // 通过 struct 读写会忽略该字段. 使用 GORM Migrator 创建表时，不会创建被忽略的字段
}

type Author struct {
	Name  string `gorm:"unique;not null;index;type:varchar(30);default:qiaocc;comment:用户名"`
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Blog{})

	// Create
	db.Create(&User{Name: "tom", Age: 29})

	// Read
	var user User
	db.First(&user, 1) // 根据整形主键查找

	// Update `gorm:"<-:create"` // 只允许读和创建
	// 更新不起作用
	db.Model(&User{}).Where("id > ?", 0).Updates(User{Name: "jerry"})

	fmt.Printf("user=%#v\n", user)
}
