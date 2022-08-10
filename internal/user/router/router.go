package router

import "zero/internal/user/app"

type Router struct {
	app *app.Application
}

func NewRouter(app *app.Application) *Router {
	return &Router{
		app: app,
	}
}
