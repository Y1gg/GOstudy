package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 全局变量
var db *sql.DB

// 连接数据库操作
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

// 创建mysql记录对应的结构体
type user struct {
	id   int
	name string
	age  int
}

func transactionDemo() {
	//1、开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed ,err:%v\n", err)
		return
	}

	//执行多条sql语句
	sqlStr1 := "update user set age=age-2 where id=1"
	sqlStr2 := "update user set age=age+2 where id=2"

	//执行sql1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql1出错,需要回滚")
		return
	}

	//执行sql2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql2出错,需要回滚")
		return
	}

	//上面两步都执行成功了，就提交本次事务
	err = tx.Commit()
	if err != nil {
		fmt.Println("提交出错了,需要回滚")
		return
	}
	fmt.Println("事务执行成功！！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	transactionDemo()

}
