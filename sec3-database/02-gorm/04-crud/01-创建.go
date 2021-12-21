package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	BirthDate time.Time `gorm:"default:null"`
	UUID      uuid.UUID
	Location  string `gorm:"type:varchar(100)"`
}

// BeforeCreate 创建钩子 BeforeSave, BeforeCreate, AfterSave, AfterCreate
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	if u.Name == "admin" {
		return errors.New("invalid role")
	}
	return
}

func main() {
	dsn := "root:123@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 记录日志
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	// Create
	//user := User{Name: "qiaocc", Age: 29, BirthDate: time.Now()}
	//result := db.Create(&user)
	//fmt.Println(user.ID)
	//fmt.Println(result.Error)        // 返回 error
	//fmt.Println(result.RowsAffected) // 返回插入记录的条数

	// 用指定的字段创建记录
	//db.Select("Name", "Age", "CreatedAt").Create(&user)
	// 创建记录, 并更新没有列出来的字段
	//db.Omit("ID", "Name", "Age", "CreatedAt").Create(&user)

	// 批量插入
	//var users = []User{
	//	{Name: "qiaocc1"},
	//	{Name: "qiaocc2"},
	//	{Name: "qiaocc3"},
	//}
	//db.Create(&users)
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}

	// 指定创建的数量, 每组多少个
	//var users2 []User
	//for i := 0; i < 1000; i++ {
	//	users2 = append(users2, User{Name: "qiaocc" + strconv.Itoa(i)})
	//}
	//db.CreateInBatches(users2, 100) // 每组100个
	//for _, user := range users2 {
	//	fmt.Println(user.ID)
	//}

	// 钩子不通过
	//db.Create(&User{Name: "admin"})
	//// 忽略钩子
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&User{Name: "admin"})

	// 通过map创建. 该方法不会填充主键等信息
	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "tom", "Age": 21,
	//})
	//// 批量
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "tom2", "Age": 21},
	//	{"Name": "tom3", "Age": 21},
	//})

	// 通过 map 创建记录
	//db.Model(User{}).Create(map[string]interface{}{
	//	"Name":     "jinzhu",
	//	"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	//})

	// 通过自定义类型创建记录
}

type Location struct {
	X, Y int
}

// Scan 方法实现了 sql.Scanner 接口
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}
