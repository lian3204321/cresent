package context

import (
	"net/http"
	"sync"
)

const (
	CharsetUTF8 = "charset=utf-8"
	ApplicationJSON = "application/json"

)



//type Context interface {
//	reset()
//	add()
//
//}
// 包含http请求的response request
type Context struct {
	Req *http.Request
	handleData map[string]interface{} //存放handle的处理结果
	handleMutex sync.Mutex   //处理handle结果的读写锁
	Resp *http.ResponseWriter
}




type HandlerFunc func(ctx *Context)