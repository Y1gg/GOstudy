package main

import (
	"errors"
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

// 事务操作
func transationDemo2() (err error) {
	//开启事务
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed,err:%v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rllback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "update user set age =23 whhere id=?"
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功")

	sqlStr1 := "select id,name,age from user where id=1"
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := "select id,name,age from user "
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed ,err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
	fmt.Println("\n====================================")
	transationDemo2()
}
