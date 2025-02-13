package services

import (
	"github.com/MohamedMosalm/To-Do-List/models"
	repositories "github.com/MohamedMosalm/To-Do-List/repositories/taskRepository"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTasksByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(taskID, userID uuid.UUID) error
	GetTaskByID(taskID uuid.UUID) (*models.Task, error)
}

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.taskRepo.CreateTask(task)
}

func (s *taskService) GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	return s.taskRepo.GetTasksByUserID(userID)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.taskRepo.UpdateTask(task)
}

func (s *taskService) DeleteTask(taskID, userID uuid.UUID) error {
	return s.taskRepo.DeleteTask(taskID, userID)
}

func (s *taskService) GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	return s.taskRepo.GetTaskByID(taskID)
}
