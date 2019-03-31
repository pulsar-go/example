package routes

import (
	"../controllers/sample"
	"../middlewares"
	"github.com/pulsar-go/pulsar/router"
)

// Register registers the application routes.
func Register() {
	router.Routes.
		Get("/", sample.Index).
		Get("/user/:name", sample.User).
		Group(&router.Options{Prefix: "/sample", Middleware: middlewares.Sample}, func(routes *router.Router) {
			routes.Get("/about", sample.About).
				Get("/about2", sample.About)
		})
}
