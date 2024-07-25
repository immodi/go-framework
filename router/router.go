package router

import (
	"fmt"
	"immmodi/framework/helpers"
	"net/http"
	"regexp"
)

type Router struct {
	routes      []RouteEntry
	middlewares []func(r *http.Request)
}

type RouteEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Response struct {
	Payload     []byte
	ContentType string
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range rtr.routes {
		match := e.Match(r)
		if !match {
			continue
		}

		for _, middleware := range rtr.middlewares {
			middleware(r)
		}

		e.Handler.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}

func (rtr *Router) Route(method, path string, handlerFunc func(r *http.Request) *Response) {
	e := RouteEntry{
		Method: method,
		Path:   path,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			response := handlerFunc(r)

			w.Header().Add("Content-Type", response.ContentType)
			w.Write(response.Payload)
		},
	}
	rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeStatic() {
	isStatic := helpers.CheckForStaticFiles()
	if isStatic {
		e := RouteEntry{
			Method: "GET",
			Path:   "/static/",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fs := http.FileServer(http.Dir("static"))
				handler := http.StripPrefix("/static/", fs)
				handler.ServeHTTP(w, r)
			},
		}
		rtr.routes = append(rtr.routes, e)
	}
}

func (re *RouteEntry) Match(r *http.Request) bool {

	if r.Method != re.Method {
		return false
	}

	if r.URL.Path != re.Path {
		re, err := regexp.Compile(`^/static/.*$`)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			println(err.Error())
		}
		if matched := re.MatchString(r.URL.Path); matched {
			return true
		}
		return false
	}

	return true
}

func (rtr *Router) UseMiddleware(middleware func(r *http.Request)) {
	rtr.middlewares = append(rtr.middlewares, middleware)
}
