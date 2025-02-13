package utils

import (
	"errors"

	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateCreateTaskDTO(dto *dtos.CreateTaskDTO) error {
	if dto.Title == "" {
		return errors.New("title is required")
	}
	if len(dto.Title) > 100 {
		return errors.New("title cannot exceed 100 characters")
	}
	if len(dto.Description) > 500 {
		return errors.New("description cannot exceed 500 characters")
	}
	return nil
}

func ExtractAndValidateUserID(c *gin.Context) (uuid.UUID, error) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		return uuid.Nil, errors.New("user ID is required")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	return userID, nil
}

func ExtractAndValidateTaskID(c *gin.Context) (uuid.UUID, error) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uuid.Nil, errors.New("invalid task ID format")
	}
	return taskID, nil
}

func ApplyTaskUpdates(existing *models.Task, updates *dtos.UpdateTaskDTO) *models.Task {
	if updates.Title != "" {
		existing.Title = updates.Title
	}
	if updates.Description != "" {
		existing.Description = updates.Description
	}
	existing.Status = updates.Status
	return existing
}
