package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 单行查询
var db *sql.DB

func initDB() (err error) {
	dsn := "root:admin@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单个记录
func queryOne(id int) {
	var u user
	sqlStr := "select id,name,age from user where id=?"

	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed , err:%v\n", err)
	}
	fmt.Printf("%#v", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	queryOne(1)
}
