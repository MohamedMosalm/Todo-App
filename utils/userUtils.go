package utils

import (
	"errors"
	"os"

	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ValidateRegisterDTO(dto *dtos.RegisterDTO) error {
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

func ValidateLoginDTO(dto *dtos.LoginDTO) error {
	if dto.Email == "" {
		return errors.New("email is required")
	}

	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func GenerateJWT(userID uuid.UUID, expiry int64) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET environment variable not set")
	}

	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     expiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
