package user

import (
	"strconv"

	"github.com/pulsar-go/example/models"
	"github.com/pulsar-go/pulsar/db"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
)

// Data ...
type Data struct {
	Users []models.User
}

// Index function
func Index(req *request.HTTP) response.HTTP {
	users := []models.User{}
	db.Builder.Take(2).All(&users)
	data := Data{Users: users}
	return response.View("users.gohtml", data)
}

// Search function
func Search(req *request.HTTP) response.HTTP {
	name := req.Params.ByName("name")
	users := []models.User{}
	db.Builder.Where("name", name).Find(&users) // Find returns more than 1
	return response.JSON(users)
}

// Store function
func Store(req *request.HTTP) response.HTTP {
	name := req.Params.ByName("name")
	age := req.Params.ByName("age") // string -> uint (32 bits)
	rand := req.Params.ByName("random")
	ageInt, err := strconv.ParseUint(age, 10, 32) // uint64 bits
	if err != nil {
		ageInt = 0
	}
	user := models.User{Name: name, Age: uint(ageInt), Random: rand}
	db.Builder.Save(&user)
	return response.JSON(user)
}

// Update function
func Update(req *request.HTTP) response.HTTP {
	id := req.Params.ByName("id")
	name := req.Params.ByName("name")
	user := models.User{}
	db.Builder.First(&user, id)
	user.Name = name
	db.Builder.Save(&user)
	return response.JSON(user)
}

// Delete ...
func Delete(req *request.HTTP) response.HTTP {
	id := req.Params.ByName("id")
	db.Builder.Where("ID", id).Delete(&models.User{})
	return response.Text("User " + id + " deleted")
}
