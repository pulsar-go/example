package models

import "github.com/pulsar-go/pulsar/db"

// User is the model representing the DB users table.
type User struct {
	db.Model
	Name   string
	Age    uint
	Random string
}
