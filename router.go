package rakhsh

import "net/http"

type Router struct {
	routes map[string]HandlerFunc
}

type HandlerFunc func(*Context)


func newRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) add(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	r.routes[key] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handler, ok := r.routes[key]; ok {
		ctx := &Context{
			Writer:  w,
			Request: req,
			Params:  map[string]string{},
		}
		handler(ctx)
		return
	}

	http.NotFound(w, req)
}
