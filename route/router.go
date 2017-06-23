package route

const (
	GET int = iota
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD


)

var RouterMethods = map[string]int{
	"GET":GET,
	"POST":POST,
	"PUT":PUT,
	"DELETE":DELETE,
	"PATCH":PATCH,
	"OPTIONS":OPTIONS,
	"HEAD":HEAD,
}

type Router interface {
	Match(method string , uri string , c *Context)
	Add(method string , pattern string ,handles HandlerFuc )
}