package dtos

import (
	"time"

	"github.com/MohamedMosalm/Todo-App/models"
	"github.com/google/uuid"
)

type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description" binding:"max=500"`
}

type UpdateTaskDTO struct {
	Title       string `json:"title" binding:"omitempty,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
	Status      bool   `json:"status"`
}

type TaskResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTaskResponseDTO(task *models.Task) *TaskResponseDTO {
	return &TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
