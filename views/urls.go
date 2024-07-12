package views

import (
	"immmodi/framework/router"
)

func GetAllRoutes() *router.Router {
	r := &router.Router{}
	r.ServeStatic()

	DefineRoutes(r)
	return r
}

func DefineRoutes(r *router.Router) {
	r.Route("GET", "/", Root)
	r.Route("GET", "/test", Test)
	r.Route("GET", "/json", Json)
	r.Route("GET", "/text", Text)
}
