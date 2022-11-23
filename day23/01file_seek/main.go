package main

import (
	"fmt"
	"io"
	"os"
)

/*
func (*File) Seek
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
Seek设置下一次读/写的位置。
offset为相对偏移量，而whence决定相对位置：
0为相对文件开头，1为相对当前位置，2为相对文件结尾。
它返回新的偏移量（相对开头）和可能的错误。
*/
func f() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./abc.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	//因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v", err)
		return
	}
	defer tmpFile.Close()

	//读取文件写入临时文件
	var ret [1]byte                //定义一个byte类型的数组
	n, err := fileObj.Read(ret[:]) //读取ret切片
	if err != nil {
		fmt.Printf("read from file failed,err:%v", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//再写入要插入的内容
	s := []byte{'c'}
	tmpFile.Write(s)
	//紧接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])

		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}

		tmpFile.Write(x[:n])
	}
	//源文件后续的也写入了临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./abc.txt")

}

func main() {
	f()
}
