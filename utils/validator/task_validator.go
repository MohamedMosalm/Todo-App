package validator

import (
	"errors"
	"github.com/MohamedMosalm/To-Do-List/dtos"
)

type TaskValidator struct{}

func NewTaskValidator() *TaskValidator {
	return &TaskValidator{}
}

func (v *TaskValidator) ValidateCreateTaskDTO(dto *dtos.CreateTaskDTO) error {
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
