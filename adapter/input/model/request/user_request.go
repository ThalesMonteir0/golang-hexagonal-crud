package request

type UserRequest struct {
	Name     string `json:"name" biding:"required"`
	Email    string `json:"email" biding:"required,email"`
	Password string `json:"password" biding:"required,min=6"`
}
