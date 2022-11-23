package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	//数据库信息
	//用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:admin@tcp(127.0.0.1:3306)/sql_test"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return

}

type user struct {
	ID   int
	Name string
	Age  int
}

func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)

	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user

	err := db.Select(&users, sqlStr)

	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err:%v\n", err)
		return
	}
	// fmt.Println("数据库连接成功")
	// sqlInjectDemo("旭旭")
	// sqlInjectDemo("xxx' or 1=1 #")
	sqlInjectDemo("xxx' union select * from user #")

}
