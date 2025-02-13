package handlers

import (
	"errors"
	"net/http"

	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/MohamedMosalm/To-Do-List/services"
	"github.com/MohamedMosalm/To-Do-List/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var createTaskDTO dtos.CreateTaskDTO

	if err := c.ShouldBindJSON(&createTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := utils.ValidateCreateTaskDTO(&createTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(createTaskDTO.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	task := models.Task{
		Title:       createTaskDTO.Title,
		Description: createTaskDTO.Description,
		Status:      createTaskDTO.Status,
		UserID:      userID,
	}

	if err := h.taskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, dtos.NewTaskResponseDTO(&task))
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID, err := utils.ExtractAndValidateUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.taskService.GetTasksByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	taskResponses := make([]dtos.TaskResponseDTO, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = *dtos.NewTaskResponseDTO(&task)
	}

	c.JSON(http.StatusOK, taskResponses)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := utils.ExtractAndValidateTaskID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.ExtractAndValidateUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateDTO dtos.UpdateTaskDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	existingTask, err := h.taskService.GetTaskByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		return
	}

	if existingTask.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this task"})
		return
	}

	updatedTask := utils.ApplyTaskUpdates(existingTask, &updateDTO)
	if err := h.taskService.UpdateTask(updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, dtos.NewTaskResponseDTO(updatedTask))
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := utils.ExtractAndValidateTaskID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.ExtractAndValidateUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTask, err := h.taskService.GetTaskByID(taskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
		return
	}

	if existingTask.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this task"})
		return
	}

	if err := h.taskService.DeleteTask(taskID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
