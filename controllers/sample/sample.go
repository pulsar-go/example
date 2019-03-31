package sample

import (
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
)

// Sample represents a sample json.
type Sample struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Index function
func Index(req *request.HTTP) response.HTTP {
	return response.View("sample.html")
}

// User function
func User(req *request.HTTP) response.HTTP {
	type data struct {
		Name string
	}
	return response.Template("sample.gohtml", data{Name: req.Params.ByName("name")})
}

// About function
func About(req *request.HTTP) response.HTTP {
	return response.JSON(Sample{Name: "Erik", Age: 22})
}
