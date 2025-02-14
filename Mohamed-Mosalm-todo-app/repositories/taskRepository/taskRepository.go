package repositories

import (
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/google/uuid"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTasksByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(taskID, userID uuid.UUID) error
	GetTaskByID(taskID uuid.UUID) (*models.Task, error)
}
