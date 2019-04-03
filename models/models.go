package models

import (
	"github.com/pulsar-go/pulsar/database"
)

func init() {
	database.AddModels(&User{})
}
