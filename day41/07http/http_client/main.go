package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//net/http Client

func main() {

	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "张旭")
	data.Set("age", "23")
	queryStr := data.Encode() //url encode之后的url
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	/*
		resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=张旭&age=18")

	*/

	//从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed ,err:", err)
		return
	}
	fmt.Println(string(b))

}
