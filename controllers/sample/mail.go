package sample

import (
	"github.com/pulsar-go/pulsar/mail"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
)

// Send sends an sample email.
func Send(req *request.HTTP) response.HTTP {
	err := mail.Create().
		To("soc@erik.cat", "ConsoleTVs@gmail.com").
		Subject("Pulsar email").
		Text("This email was sent using pulsar.").
		Send()
	if err != nil {
		return response.Text("Error: " + err.Error())
	}
	return response.Text("Email sent.")
}
