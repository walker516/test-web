package router

import (
	"backend/internal/handler"
	"backend/internal/interceptor/authmw"
	"backend/internal/interceptor/corsmw"
	"backend/internal/repository"
	"backend/internal/usecase"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Router struct {
	echo       *echo.Echo
	myDB       *sqlx.DB
	handlers   *handler.Handlers
	middleware *Middleware
}

type Middleware struct {
	Auth echo.MiddlewareFunc
	CORS echo.MiddlewareFunc
}

func NewRouter(e *echo.Echo, db *sqlx.DB) *Router {
	return &Router{
		echo: e,
		myDB: db,
	}
}

func (r *Router) SetupRoutes() {
	r.setupMiddleware()
	r.setupHandlers()
	r.registerRoutes()
}

func (r *Router) setupMiddleware() {
	r.middleware = &Middleware{
		Auth: authmw.RequireUserID(),
		CORS: corsmw.NewCORSConfig(),
	}
	r.echo.Use(r.middleware.CORS)
}

func (r *Router) setupHandlers() {
	repos := repository.NewRepositories(r.myDB)
	services := usecase.NewUsecases(repos)
	r.handlers = handler.NewHandlers(services)
}

func (r *Router) registerRoutes() {
	// ✅ CORS のプリフライトリクエスト (OPTIONS) を明示的に登録
	r.echo.OPTIONS("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	userAPI := r.echo.Group("/api/user/v1")

	userAPI.OPTIONS("/users", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	}) 

	userAPI.GET("/users", r.handlers.User.GetAll)
	userAPI.GET("/users/:id", r.handlers.User.GetByID)
	userAPI.POST("/users", r.handlers.User.Create)
	userAPI.PUT("/users/:id", r.handlers.User.Update)
	userAPI.DELETE("/users/:id", r.handlers.User.Delete)
}

