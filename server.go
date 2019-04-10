package main

import (
	"log"

	"github.com/pulsar-go/example/routes"
	"github.com/pulsar-go/pulsar"
)

func main() {
	// Get the settings from the configuration files.
	// config.Set("./server.toml")
	// Set the application routes.
	routes.Register()
	// Serve the HTTP server.
	log.Fatalln(pulsar.Serve())
}
