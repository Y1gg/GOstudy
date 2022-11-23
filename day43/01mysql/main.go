package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 没有运行出来
// 定义一个全局变量db
var db *sql.DB

func initDB() (err error) {
	//数据库信息
	//用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:admin@tcp(127.0.0.1:3306)/sql_test"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                  //dsn格式不正确的时候会报错
		return err
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// go连接 mysql 示例
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功!")

	var u1 user
	//1.写查询单挑记录的sql语句
	sqlStr := "select id,name,age form user where id=?"
	//2.执行
	//从连接池里拿一个连接出来去数据库查询单条记录
	rowObj := db.QueryRow(sqlStr, 1)
	//3.拿到结果
	err = rowObj.Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	//打印结果
	fmt.Printf("u1:%#v\n", u1)

}
