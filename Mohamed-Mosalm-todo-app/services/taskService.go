package services

import (
	"github.com/MohamedMosalm/Todo-App/models"
	taskRepository "github.com/MohamedMosalm/Todo-App/repositories/taskRepository"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTasksByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error
	DeleteTask(taskID, userID uuid.UUID) error
	GetTaskByID(taskID uuid.UUID) (*models.Task, error)
}

type taskService struct {
	taskRepo taskRepository.TaskRepository
}

func NewTaskService(taskRepo taskRepository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.taskRepo.CreateTask(task)
}

func (s *taskService) GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	return s.taskRepo.GetTasksByUserID(userID)
}

func (s *taskService) UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error {
	return s.taskRepo.UpdateTask(taskID, updates)
}

func (s *taskService) DeleteTask(taskID, userID uuid.UUID) error {
	return s.taskRepo.DeleteTask(taskID, userID)
}

func (s *taskService) GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	return s.taskRepo.GetTaskByID(taskID)
}
