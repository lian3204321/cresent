package route

import (
	"cresent/context"
)

const (
	GET        int = iota
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD
	METHOD_LEN
)

var RouterMethods = map[string]int{
	"GET":     GET,
	"POST":    POST,
	"PUT":     PUT,
	"DELETE":  DELETE,
	"PATCH":   PATCH,
	"OPTIONS": OPTIONS,
	"HEAD":    HEAD,
}

type Router interface {
	Match(method string, uri string) ([]context.HandlerFunc, string)
	Add(method string, pattern string, handles []context.HandlerFunc) string
}
type RouteNode interface {
	Name(name string)
}

//先匹配节点
type Tree struct {
	nodes   []*Leaf               //子节点,下标对应的method
	handles []context.HandlerFunc //匹配后的调用方法，在路由完全匹配成功采取先进先出队列的方式调用
}

type Leaf struct {
	pathName string           //路由 如 '/' ,'avc/'
	parent   *Leaf            //父级节点 可以为null
	nodes    map[string]*Leaf // 子节点
	handles  []context.HandlerFunc
	nodeName string //节点名称对应上级目录中的
}

type Node struct {

}

func (tr *Tree) Match(method string, uri string)([]context.HandlerFunc, string ){

}
func (tr *Tree) Add(method string, pattern string, handles []context.HandlerFunc) *Leaf {
	//获取路由节点
}

func (tr *Tree) get(nodeName string) {

}
