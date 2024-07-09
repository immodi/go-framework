package routes

import "immmodi/framework/handlers"

func GetAllRoutes() *handlers.Router {
	r := &handlers.Router{}
	DefineRoutes(r)
	return r
}

func DefineRoutes(r *handlers.Router) {
	r.Route("GET", "/", Root)
	r.Route("GET", "/test", Test)
	r.Route("GET", "/j", J)
}
