package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	CompanyRefer int
	CreditCard   CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	dsn := "root:123@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("db open err:%v\n\n", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&User{}, &CreditCard{})
	if err != nil {
		fmt.Printf("auto migrate failed, err:%v\n", err)
	}

}
