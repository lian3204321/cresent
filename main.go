package main

import (
	"net/http"
	"io"
	"fmt"
)

func hello(rw http.ResponseWriter, rq *http.Request) {
	io.WriteString(rw, "hello world")
}

// handle
type moonHandle struct {
	routerMap map[string]interface{} //path map   end with requestHandle

}

func (moonHandle) RegisterRouter() {

}



func (moonHandle) ServeHTTP(rw http.ResponseWriter, pr *http.Request) {
	fmt.Println(pr.RequestURI, pr.Host, pr.URL.Path , pr.Method)

	//router
	if pr.URL.Path == "/"{

	}

	if pr.URL.Path == "/aaa"{

	}

}

type moon struct {
}

func (moon) Run(addr string) {
	//interface struct
	moonh := &moonHandle{}
	http.ListenAndServe(addr, moonh)
}

func main() {
	//简单实现
	//http.HandleFunc("/" , hello)

	//http.ListenAndServe(":8081" ,nil)

	moonRoute := &moon{}
	moonRoute.Run(":8081")
}
