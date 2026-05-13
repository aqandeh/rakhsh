package rakhsh

import "net/http"

type App struct {
	router *Router
}

func New() *App {
	return &App{
		router: newRouter(),
	}
}


func (a *App) GET(path string, handler HandlerFunc) {
	a.router.add("GET", path, handler)
}

func (a *App) POST(path string, handler HandlerFunc) {
	a.router.add("POST", path, handler)
}

func (a *App) Listen(addr string) error {
	return http.ListenAndServe(addr, a.router)
}
