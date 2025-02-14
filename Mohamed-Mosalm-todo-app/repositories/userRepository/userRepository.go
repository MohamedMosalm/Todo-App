package repositories

import (
	"github.com/MohamedMosalm/Todo-App/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
}
