package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"zero/internal/auth/app"
)

type Router struct {
	app *app.Application
}

func NewRouter(params app.ApplicationParams) *Router {
	wg := &sync.WaitGroup{}
	a := app.MustNewApplication(context.TODO(), wg, params)
	return &Router{
		app: a,
	}
}

// Load initializes the routing of the application.
func (r *Router) Load() http.Handler {
	router := gin.New()

	setGeneralMiddlewares(context.TODO(), router)
	registerAPIHandlers(router, r.app)

	return router
}

func registerAPIHandlers(router *gin.Engine, app *app.Application) {
	// Build middlewares
	//BearerToken := NewAuthMiddlewareBearer(app)

	// We mount all handlers under /api path
	r := router.Group("/api")
	v1 := r.Group("/v1")

	// Add health-check
	v1.GET("/health", handlerHealthCheck())

	// Add auth namespace
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/users", RegisterUser(app))
		authGroup.POST("/users/login", LoginUser(app))
	}
}

// SetGeneralMiddlewares add general-purpose middlewares
func setGeneralMiddlewares(ctx context.Context, ginRouter *gin.Engine) {
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(CORSMiddleware())
	ginRouter.Use(requestid.New())
	ginRouter.Use(LoggerMiddleware(ctx))
}
