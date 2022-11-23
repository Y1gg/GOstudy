package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := `<h1 style="color:blue">你好！张旭旭旭旭旭旭</h1>`
	w.Write([]byte(str))
}
func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	//对于GET请求，参数都放在URL上(query param),请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

// net/http  server
func main() {
	http.HandleFunc("/posts/Go/15_sockrt/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
