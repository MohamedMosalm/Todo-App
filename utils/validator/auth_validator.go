package validator

import (
	"errors"
	"github.com/MohamedMosalm/To-Do-List/dtos"
)

type AuthValidator struct{}

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{}
}

func (v *AuthValidator) ValidateRegisterDTO(dto *dtos.RegisterDTO) error {
	if dto.FirstName == "" {
		return errors.New("first name is required")
	}
	if dto.LastName == "" {
		return errors.New("last name is required")
	}
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (v *AuthValidator) ValidateLoginDTO(dto *dtos.LoginDTO) error {
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
