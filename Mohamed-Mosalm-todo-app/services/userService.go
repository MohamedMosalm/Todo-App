package services

import (
	"github.com/MohamedMosalm/Todo-App/models"
	userRepository "github.com/MohamedMosalm/Todo-App/repositories/userRepository"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
}

type userService struct {
	userRepo userRepository.UserRepository
}

func NewUserService(userRepo userRepository.UserRepository) UserService {
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
