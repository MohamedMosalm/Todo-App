package services

import (
	"github.com/MohamedMosalm/To-Do-List/models"
	repositories "github.com/MohamedMosalm/To-Do-List/repositories/userRepository"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) FindUserByEmail(email string) (*models.User, error) {
	return s.userRepo.FindUserByEmail(email)
}

func (s *userService) FindUserByID(id uuid.UUID) (*models.User, error) {
	return s.userRepo.FindUserByID(id)
}
