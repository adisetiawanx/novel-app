package request

type AuthRegisterRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8,max=32"`
	RePassword string `json:"rePassword" validate:"required,min=8,max=32,eqfield=Password"`
	Phone      string `json:"phone" validate:"required,numeric"`
}
