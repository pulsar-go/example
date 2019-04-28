package main

import (
	"log"

	"github.com/pulsar-go/example/routes"
	"github.com/pulsar-go/pulsar"
)

func main() {
	// Set the application routes.
	routes.Register()
	// Serve the HTTP server.
	log.Fatalln(pulsar.Serve())
}
