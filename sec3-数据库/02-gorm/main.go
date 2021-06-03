package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("db open err:%v\n\n", err)
	}

	// 自动迁移
	//err = db.AutoMigrate(&UserInfo{})
	//if err != nil {
	//	fmt.Printf("auto migrate failed, err:%v\n", err)
	//}
	//
	//// 增
	//db.Create(&UserInfo{
	//	Name:   "marry",
	//	Gender: "女",
	//	Hobby:  "旅游",
	//})

	// 查
	var u UserInfo
	db.First(&u, 1) // 主键=1
	//fmt.Printf("id=%v, name=%v, gender=%v, hobby:%v\n", u.ID, u.Name, u.Gender, u.Hobby)
	//db.First(&u, "gender=?", "男") // 性别男
	//fmt.Printf("id=%v, name=%v, gender=%v, hobby:%v\n", u.ID, u.Name, u.Gender, u.Hobby)

	// 改
	//db.Model(&u).Update("gender", "女")
	//db.Model(&u).Updates(UserInfo{Gender: "男", Name: "qiaocc"})
	//db.Model(&u).Updates(map[string]interface{}{"Gender": "女", "Name": "Marry",
	//})

	// 删
	db.Delete(&u, 1)
}
