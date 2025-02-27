package repositories

import (
	"github.com/MohamedMosalm/Todo-App/models"
	"github.com/google/uuid"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTasksByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error
	DeleteTask(taskID, userID uuid.UUID) error
	GetTaskByID(taskID uuid.UUID) (*models.Task, error)
}
