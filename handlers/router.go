package handlers

import "net/http"

type Router struct {
	routes []RouteEntry
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

func (re *RouteEntry) Match(r *http.Request) bool {
	if r.Method != re.Method {
		return false // Method mismatch
	}

	if r.URL.Path != re.Path {
		return false // Path mismatch
	}

	return true
}
