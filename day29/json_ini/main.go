package main

import (
	"errors"
	"fmt"
	"reflect"
)

//ini配置文件解析器

// Mysqlconfig Mysql配置结构体
type Mysqlconfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//一、参数的校验
	//1.1传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //新创建一个错误
		return
	}
	//1.2传进来的data参数必须是结构体类型指针(因为配置文件中各种键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") //新创建一个错误
		return
	}
	//二、读文件的得到字节类型数据
	//三、一行一行得到数据
	//3.1如果是注释就跳过
	//3.2如果是[开头就表示是节(section)
	//3.3如果不是[开头就是=分割的键值对
	return
}

func main() {
	var mc Mysqlconfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

}
