package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:123@tcp(127.0.0.1:3306)/go"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 单行查询
func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:=%v\n", err)
		return
	}
	fmt.Printf("id=%v, name=%v, age=%v", u.id, u.name, u.age)
}

func insertRow() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "jerry", 25)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %v\n", theID)

}

func queryMultiRow() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v", err)
		return
	}

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id=%v, name=%v, age=%v\n", u.id, u.name, u.age)
	}
}

func updateRow() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 25, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRow() {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
	}
	fmt.Printf("delete success, affected rows:%v\n", n)
}

func prepareQuery() {
	sqlStr := "select id, name, age from user where id>?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id=%v, name=%v, age=%v\n", u.id, u.name, u.age)
	}
}

func transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
	}
	sqlStr1 := "update user set age = 30 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql1 failed, err:%v\n")
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v", err)
		return
	}

	// 单行查询
	//queryRow()

	// 多行查询
	//queryMultiRow()

	// 插入数据
	//insertRow()

	// 修改数据
	//updateRow()

	// 删除数据
	//deleteRow()

	// 预处理
	//prepareQuery()

	// 事务
	transaction()
}
