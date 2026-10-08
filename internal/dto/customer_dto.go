package dto

type CreateCustomerRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,min=10,max=15"`
	Address string `json:"address" validate:"required,min=5,max=200"`	
}
