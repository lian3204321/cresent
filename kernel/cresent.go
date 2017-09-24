package kernel

import (
	"cresent/route"
	"cresent/context"
	"net/http"
	"fmt"
)

const (
	DEV  = "develop"
	TEST = "test"
	PROD = "production"
)

// 环境变量
var ENV string

type Cresent struct {
	debug bool
	name  string
	router route.Router
}

func init() {
	ENV = DEV //todo 做环境变量的判断
	appInstances = make(map[string]*Cresent)
}

var appInstances map[string]*Cresent

const APP_NAME = "cresent"

//设置单例模式
func Instance(name string) *Cresent {
	if name == "" {
		name = APP_NAME
	}
	if appInstances[name] == nil {
		appInstances[name] = New()
		appInstances[name].name = name
	}
	return appInstances[name]
}

func Default() *Cresent {
	return Instance(APP_NAME)
}

func New() *Cresent {
	c := new(Cresent)
	return c
}

// 设置http服务
func (c *Cresent) Server(addr string) *http.Server {
	//todo http server的配置需要添加
	return &http.Server{Addr: addr}
}

func (c *Cresent) Run(addr string) {
	c.run(c.Server(addr))
}

// 启动服务
func (c *Cresent) run(s *http.Server, files ...string) {
	s.Handler = c
	if (len(files) == 0) {
		s.ListenAndServe()
	}
}

func (c *Cresent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//进行请求和响应数据的封装
	con := new(context.Context) //todo 需要加入pool

	//c.Route()
	//进行路由的匹配
	uri := r.RequestURI
	fmt.Println(uri)
	//进行数据的处理

	//结束请求
}

func (c *Cresent) Route() route.Router {
	//todo 需要做单例和di
	//c.router = new(route.Tree)
	return c.router
}
