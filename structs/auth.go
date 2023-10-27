package structs

type SignUpInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpOutput struct {
	Message string `json:"message"`
}
