package router

import "zero/internal/user/app"

type Router struct {
	app *app.Application
}

func NewRouter() *Router {
	a := new(app.Application)
	return &Router{
		app: a,
	}
}
