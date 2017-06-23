package kernel

import "cresent/route"

const (
	DEV = "develop"
	TEST = "test"
	PROD = "production"
)


var ENV string

type Cresent struct {
	debug  bool
	name string
	router route.Router
}

type HandelFunc func(*Context)