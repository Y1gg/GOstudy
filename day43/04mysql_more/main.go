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
	fmt.Printf("%#v\n", u)
}
func querMultiRowDemo(id int) {
	sqlStr := "select id,name,age from user where id>?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed ,err:%v\n", err)
		return
	}
	//关闭rows释放持有的数据库链接

	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("sacn failed ,err:%v\n", err)
			return
		}
		fmt.Printf("%#v\n", u)
	}
}

func insert(newName string, newAge int) {
	//1、写sql语句
	sqlStr := `insert into user(name,age) values(?,?)`
	ret, err := db.Exec(sqlStr, newName, newAge)
	if err != nil {
		fmt.Printf("insert failed ,err:%v\n", err)
		return
	}

	//如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

func updateRowDemo(newAge, id int) {
	sqlStr := "update user set age =? where id=?"
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed ,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success ,affected rows:%d\n", n)
}

//删除数据

func deleteRowDemo(id int) {
	sqlStr := "delete  from user  where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed ,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows:%d\n", n)
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name,age ) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("许斌", 23)
	if err != nil {
		fmt.Printf("insert failed ,err:%v\n", err)
		return
	}
	_, err = stmt.Exec("陆小栎", 23)
	if err != nil {
		fmt.Printf("insert failed ,err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed ,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	/*
		fmt.Println("\n====================================查询单条")
		queryOne(1)
		queryOne(2)
		fmt.Println("\n====================================查询多条")
		querMultiRowDemo(2)
		fmt.Println("\n====================================添加数据")
		insert("张孟宇", 24)
		fmt.Println("\n====================================更新数据")
		updateRowDemo(22, 2)
		deleteRowDemo(3)
	*/
	prepareInsertDemo()
}
