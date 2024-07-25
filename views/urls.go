package views

import (
	"immmodi/framework/middlewares"
	"immmodi/framework/router"
)

func Router() *router.Router {
	rtr := &router.Router{}
	rtr.ServeStatic()

	rtr.UseMiddleware(middlewares.Logger)
	DefineRoutes(rtr)
	return rtr
}

func DefineRoutes(r *router.Router) {
	r.Route("GET", "/", Root)
	r.Route("GET", "/test", Test)
	r.Route("GET", "/json", Json)
	r.Route("GET", "/text", Text)
	r.Route("GET", "/again", BasicEdit)
	r.Route("POST", "/again", BasicEditPost)
}
