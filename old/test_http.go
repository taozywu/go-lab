package main

import (
	"html/template"
	// "log"
	"net/http"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//加载模板
		tpl, _ := template.ParseFiles("demo.html")
		//执行模板
		tpl.Execute(w, Person{Name: "test", Age: 18})
	})
	server := http.Server{
		Addr:        ":9090",         //监听地址和端口
		Handler:     mux,             //Handle
		ReadTimeout: 5 * time.Second, //读取超时5s
	}
	//启动监听
	server.ListenAndServe()
}
