package models

import "github.com/pulsar-go/pulsar/database"

// Register registers the models to the database.
func Register() {
	database.AddModels(&User{})
}
