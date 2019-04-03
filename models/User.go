package models

import "github.com/pulsar-go/pulsar/database"

// User is the model representing the DB users table.
type User struct {
	database.Model
	Name string
	Age  uint
}
