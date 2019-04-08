package routes

import (
	"github.com/pulsar-go/example/controllers/sample"
	"github.com/pulsar-go/example/controllers/user"
	"github.com/pulsar-go/example/middlewares"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
	"github.com/pulsar-go/pulsar/router"
)

// Register registers the application routes.
func Register() {
	router.Routes.
		Get("/", func(req *request.HTTP) response.HTTP {
			return response.Text("Accept-Encoding: " + req.Request.Header.Get("Accept-Encoding"))
		}).
		Get("/email", sample.Send).
		Get("/job", user.SampleJob).
		Group(&router.Options{Prefix: "/user", Middleware: middlewares.Sample}, func(routes *router.Router) {
			routes.Get("/", user.Index).
				Get("/create/:name/:age/:random", user.Store).
				Get("/update/:id/:name", user.Update).
				Get("/delete/:id", user.Delete).
				Get("/search/:name", user.Search)
		})
}
