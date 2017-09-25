package main

import (
	"net/http"
	"io"
	"fmt"
	"strings"
)

func hello(rw http.ResponseWriter, rq *http.Request) {
	io.WriteString(rw, "hello world")
}

//define http method
const (
	GET, POST, PUT, DELETE = "GET", "POST", "PUT", "DELETE"
)

var METHOD_MAP = []string{GET, POST, PUT, DELETE}

// handle
type moonHandle struct {
	routerMap map[string]interface{} //path map   end with requestHandle
}

func (*moonHandle) RegisterRouter(methods []string, path string, controller func(rw http.ResponseWriter, pr *http.Request)) {
	for _, v := range methods {
		for _,vMehtod := range METHOD_MAP{
			if strings.ToUpper(v) == vMehtod{

			}
		}
	}
}

func routePath(rw http.ResponseWriter, pr *http.Request) {
	io.WriteString(rw, "xman path")
}

func urlRoutePath(rw http.ResponseWriter, pr *http.Request) {
	io.WriteString(rw, "url aaa xman path")
}

func (*moonHandle) match(realPath []string) func(rw http.ResponseWriter, pr *http.Request) {

}

func (moonHandle) ServeHTTP(rw http.ResponseWriter, pr *http.Request) {
	//fmt.Println(pr.RequestURI, pr.Host, pr.URL.Path, pr.Method)

	paths := strings.Split(pr.URL.Path, "/")
	fmt.Println(paths)
	var firstPath []string
	firstPath[0] = "/"
	realPath := append(firstPath, paths ...)

	//router
	if pr.URL.Path == "/" {
		routePath(rw, pr)
	}

	if pr.URL.Path == "/aaa" {
		urlRoutePath(rw, pr)
	}

}

type moon struct {
}

func (moon) Run(addr string) {
	//interface struct
	moonh := &moonHandle{}
	moonh.RegisterRouter([]string{"GET", "POST"}, "/", routePath)
	moonh.RegisterRouter([]string{"GET", "POST"}, "/aaaaa", urlRoutePath)

	http.ListenAndServe(addr, moonh)
}

func main() {
	//简单实现
	//http.HandleFunc("/" , hello)

	//http.ListenAndServe(":8081" ,nil)

	moonRoute := &moon{}
	moonRoute.Run(":8081")
}
