package main

import (
	"log"

	"./controllers/sample"
	"./middlewares"
	"github.com/pulsar-go/pulsar"
	"github.com/pulsar-go/pulsar/config"
	"github.com/pulsar-go/pulsar/router"
)

func main() {
	// Get the settings from the configuration files.
	config.Set("./server.toml")
	// Set the application routes.
	routes := router.Create()
	routes.Get("/", sample.Index)
	routes.Get("/user/:name", sample.User)
	routes.Group(&router.Options{Prefix: "/sample", Middleware: middlewares.Middle}, func(routes *router.Router) {
		routes.Get("/about", sample.About)
	})
	// Serve the HTTP server.
	log.Fatalln(pulsar.Serve(routes))
}
