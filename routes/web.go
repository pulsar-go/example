package routes

import (
	"github.com/pulsar-go/example/controllers/user"
	"github.com/pulsar-go/example/middlewares"
	"github.com/pulsar-go/pulsar/router"
)

// Register registers the application routes.
func Register() {
	router.Routes.Group(&router.Options{Prefix: "/user", Middleware: middlewares.Sample}, func(routes *router.Router) {
		routes.Get("/", user.Index).
			Get("/create/:name/:age", user.Store).
			Get("/update/:id/:name", user.Update).
			Get("/delete/:id", user.Delete).
			Get("/search/:name", user.Search)
	})
}
