package request

import (

"github.com/asaskevich/govalidator"
rest "github.com/cheebo/gorest"
)

type CreateUserRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,printableascii"`
	Name     string `json:"name" valid:"required,printableascii"`
}

func (r CreateUserRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return rest.ValidateResponse(err)
}
