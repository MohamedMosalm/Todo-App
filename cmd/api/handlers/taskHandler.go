package handlers

import (
	"net/http"

	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/MohamedMosalm/To-Do-List/services"
	"github.com/MohamedMosalm/To-Do-List/utils/errors"
	"github.com/MohamedMosalm/To-Do-List/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var createTaskDTO dtos.CreateTaskDTO

	if err := c.ShouldBindJSON(&createTaskDTO); err != nil {
		appErr := errors.ErrInvalidRequest
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		appErr := errors.ErrInvalidUserID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	task := models.Task{
		Title:       createTaskDTO.Title,
		Description: createTaskDTO.Description,
		Status:      false,
		UserID:      userID,
	}

	if err := h.taskService.CreateTask(&task); err != nil {
		appErr := errors.ErrCreateTaskFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	httputil.SendSuccess(c, http.StatusCreated, "Task created successfully", dtos.NewTaskResponseDTO(&task))
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		appErr := errors.ErrInvalidUserID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	tasks, err := h.taskService.GetTasksByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			httputil.SendSuccess(c, http.StatusOK, "No tasks found", []dtos.TaskResponseDTO{})
			return
		}
		appErr := errors.ErrFetchTasksFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	taskResponses := make([]dtos.TaskResponseDTO, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = *dtos.NewTaskResponseDTO(&task)
	}

	httputil.SendSuccess(c, http.StatusOK, "Tasks retrieved successfully", taskResponses)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		appErr := errors.ErrInvalidTaskID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		appErr := errors.ErrInvalidUserID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	var updateDTO dtos.UpdateTaskDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		appErr := errors.ErrInvalidRequest
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	existingTask, err := h.taskService.GetTaskByID(taskID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			httputil.HandleError(c, errors.ErrTaskNotFound)
			return
		}
		appErr := errors.ErrFetchTasksFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	if existingTask.UserID != userID {
		httputil.HandleError(c, errors.ErrUnauthorized)
		return
	}

	existingTask.Title = updateDTO.Title
	existingTask.Description = updateDTO.Description
	existingTask.Status = updateDTO.Status

	if err := h.taskService.UpdateTask(existingTask); err != nil {
		appErr := errors.ErrUpdateTaskFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, "Task updated successfully", dtos.NewTaskResponseDTO(existingTask))
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		appErr := errors.ErrInvalidTaskID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		appErr := errors.ErrInvalidUserID
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	existingTask, err := h.taskService.GetTaskByID(taskID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			httputil.HandleError(c, errors.ErrTaskNotFound)
			return
		}
		appErr := errors.ErrFetchTasksFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	if existingTask.UserID != userID {
		httputil.HandleError(c, errors.ErrUnauthorized)
		return
	}

	if err := h.taskService.DeleteTask(taskID, userID); err != nil {
		appErr := errors.ErrDeleteTaskFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, "Task deleted successfully", nil)
}
