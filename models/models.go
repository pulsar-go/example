package models

import (
	"github.com/pulsar-go/pulsar/db"
)

func init() {
	db.AddModels(&User{})
}
