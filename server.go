package main

import (
	"log"

	"./routes"
	"github.com/pulsar-go/pulsar"
	"github.com/pulsar-go/pulsar/config"
)

func main() {
	// Get the settings from the configuration files.
	config.Set("./server.toml")
	// Set the application routes.
	routes.Register()
	// Serve the HTTP server.
	log.Fatalln(pulsar.Serve())
}
