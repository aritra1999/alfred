package tests

import (
	service "albert/services"
	. "albert/structs"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("SignUp Service", func() {
		g.Describe("Happy Case", func() {
			g.It("Creates a new user", func() {
				input := SignUpInput{
					Email:    "email",
					Password: "password",
				}

				response := service.SignUp(input)
				expected := SignUpOutput{
					Message: "User created successfully",
				}

				g.Assert(response).Equal(expected)
			})
		})
	})
}
