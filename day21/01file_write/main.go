package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
func m1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err%v", err)
		return
	}

	//write
	fileObj.Write([]byte("你说你最爱丁香花\n"))
	//writeString
	fileObj.WriteString("I love MingYue")
	fileObj.Close()
}

func m2() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err%v", err)
		return
	}

	defer fileObj.Close()

	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("亚拉索 ，那就是青藏高原~~~\n") //写到缓存中
	wr.Flush()                          //将缓存写到文件

}

func m3() {
	str := "既然选择了远方，便只顾风雨兼程"
	err := ioutil.WriteFile("./x.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err", err)
		return
	}
}

func main() {
	m1()
	m2()
	m3()
}
