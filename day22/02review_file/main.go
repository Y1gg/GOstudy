package main

import (
	"fmt"
	"os"
)

//文件操作

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	// defer fileObj.Close()
	// 不能放在上面↑，因为当err有值的时候，fileObj就是nil，nil不能调用Close
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	defer fileObj.Close()

}
func main() {
	f1()
}
