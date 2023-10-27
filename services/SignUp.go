package service

import (
	"albert/models"
	. "albert/structs"
)

func SignUp(input SignUpInput) SignUpOutput {
	var response SignUpOutput

	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	response.Message = "User created successfully"

	return response
}
