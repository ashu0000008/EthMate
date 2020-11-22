package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello) //设置访问路由
	//创建监听端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var remote = r.RemoteAddr
	fmt.Fprintf(w, "Hello "+remote)
}
