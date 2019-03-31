package middlewares

import (
	"log"

	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
	"github.com/pulsar-go/pulsar/router"
)

func Middle(next router.Handler) router.Handler {
	return router.Handler(func(req *request.HTTP) response.HTTP {
		log.Println("Before route middleware")
		r := next(req)
		log.Println("After route middleware")
		return r
	})
}
