package routes

import (
	"log"
	"net/http"

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
			req.Writer.Header().Set("Server", "Pulsar")
			return response.Text("Accept-Encoding: " + req.Request.Header.Get("Accept-Encoding"))
		}).
		Post("/cors", func(req *request.HTTP) response.HTTP {
			return response.Text("CORS ok")
		}).
		Get("/status", func(req *request.HTTP) response.HTTP {
			return response.JSONWithCode(map[string]string{
				"error": "Some sample error",
			}, http.StatusBadRequest)
		}).
		Post("/json", func(req *request.HTTP) response.HTTP {
			data := struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{}
			if err := req.JSON(&data); err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			return response.JSON(data)
		}).
		Get("/static", func(req *request.HTTP) response.HTTP {
			return response.Static("sample")
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
