package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MohamedMosalm/To-Do-List/config"
	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/MohamedMosalm/To-Do-List/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) FindUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*models.User), args.Error(1)
}

func (m *MockUserService) FindUserByID(id uuid.UUID) (*models.User, error) {
	args := m.Called(id)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*models.User), args.Error(1)
}

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskService) GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error {
	args := m.Called(taskID, updates)
	return args.Error(0)
}

func (m *MockTaskService) DeleteTask(taskID, userID uuid.UUID) error {
	args := m.Called(taskID, userID)
	return args.Error(0)
}

func (m *MockTaskService) GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	args := m.Called(taskID)
	return args.Get(0).(*models.Task), args.Error(1)
}

func setupUserRouter(authHandler *AuthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/auth/register", authHandler.Register)
	router.POST("/api/auth/login", authHandler.Login)
	return router
}

func setupTaskRouter(taskHandler *TaskHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/tasks", taskHandler.CreateTask)
	router.GET("/api/tasks", taskHandler.GetTasks)
	router.PUT("/api/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/api/tasks/:id", taskHandler.DeleteTask)
	return router
}

func TestRegister(t *testing.T) {
	mockUserService := new(MockUserService)
	jwtService, _ := auth.NewJWTService("test_secret")
	passwordService := auth.NewPasswordService()
	authHandler := &AuthHandler{
		userService:     mockUserService,
		jwtService:      jwtService,
		passwordService: passwordService,
	}

	router := setupUserRouter(authHandler)

	registerDTO := dtos.RegisterDTO{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "1234567890",
		Password:  "password123",
	}

	mockUserService.On("FindUserByEmail", registerDTO.Email).Return(nil, nil)
	mockUserService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	body, _ := json.Marshal(registerDTO)
	req, _ := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUserService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockUserService := new(MockUserService)
	passwordService := auth.NewPasswordService()

	jwtService, err := auth.NewJWTService("test_secret")
	assert.NoError(t, err)

	authHandler := &AuthHandler{
		userService:     mockUserService,
		jwtService:      jwtService,
		passwordService: passwordService,
	}

	router := setupUserRouter(authHandler)

	loginDTO := dtos.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	hashedPassword, err := passwordService.HashPassword(loginDTO.Password)
	assert.NoError(t, err)

	testUser := &models.User{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     loginDTO.Email,
		Password:  hashedPassword,
	}

	mockUserService.On("FindUserByEmail", loginDTO.Email).Return(testUser, nil)

	body, err := json.Marshal(loginDTO)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	t.Logf("Response body: %v", response)

	data := response["data"].(map[string]interface{})
	assert.NotEmpty(t, data["access_token"])
	assert.Equal(t, "Bearer", data["token_type"])

	mockUserService.AssertExpectations(t)
}

func TestCreateTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	taskHandler := NewTaskHandler(mockTaskService, config.AppConfig{})

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		userID := uuid.New()
		c.Set("user_id", userID.String())
		c.Next()
	})

	router.POST("/api/tasks", taskHandler.CreateTask)

	mockTaskService.On("CreateTask", mock.AnythingOfType("*models.Task")).Return(nil)

	createTaskDTO := dtos.CreateTaskDTO{
		Title:       "Test Task",
		Description: "Test Description",
	}

	body, _ := json.Marshal(createTaskDTO)
	req, _ := http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockTaskService.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockTaskService := new(MockTaskService)
	taskHandler := NewTaskHandler(mockTaskService, config.AppConfig{})

	router := gin.Default()

	userID := uuid.New()

	router.Use(func(c *gin.Context) {
		c.Set("user_id", userID.String())
		c.Next()
	})

	router.GET("/api/tasks", taskHandler.GetTasks)

	tasks := []models.Task{
		{
			ID:          uuid.New(),
			Title:       "Task 1",
			Description: "Description 1",
			Status:      false,
			UserID:      userID,
		},
		{
			ID:          uuid.New(),
			Title:       "Task 2",
			Description: "Description 2",
			Status:      true,
			UserID:      userID,
		},
	}

	mockTaskService.On("GetTasksByUserID", userID).Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/tasks", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	data, ok := response["data"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, data, 2)

	mockTaskService.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	taskHandler := NewTaskHandler(mockTaskService, config.AppConfig{})

	router := gin.Default()

	userID := uuid.New()
	taskID := uuid.New()

	router.Use(func(c *gin.Context) {
		c.Set("user_id", userID.String())
		c.Next()
	})

	router.PUT("/api/tasks/:id", taskHandler.UpdateTask)

	existingTask := &models.Task{
		ID:          taskID,
		Title:       "Old Task",
		Description: "Old Description",
		Status:      false,
		UserID:      userID,
	}

	mockTaskService.On("GetTaskByID", taskID).Return(existingTask, nil)
	mockTaskService.On("UpdateTask", taskID, mock.AnythingOfType("map[string]interface {}")).Run(func(args mock.Arguments) {
		updates := args.Get(1).(map[string]interface{})
		if title, ok := updates["title"].(string); ok {
			existingTask.Title = title
		}
		if description, ok := updates["description"].(string); ok {
			existingTask.Description = description
		}
		if status, ok := updates["status"].(bool); ok {
			existingTask.Status = status
		}
	}).Return(nil)

	updateTaskDTO := dtos.UpdateTaskDTO{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      true,
	}

	body, _ := json.Marshal(updateTaskDTO)
	req, _ := http.NewRequest(http.MethodPut, "/api/tasks/"+taskID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Updated Task", data["title"])
	assert.Equal(t, "Updated Description", data["description"])
	assert.Equal(t, true, data["status"])

	mockTaskService.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	taskHandler := NewTaskHandler(mockTaskService, config.AppConfig{})

	router := gin.Default()

	userID := uuid.New()
	taskID := uuid.New()

	router.Use(func(c *gin.Context) {
		c.Set("user_id", userID.String())
		c.Next()
	})

	router.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	existingTask := &models.Task{
		ID:          taskID,
		Title:       "Task to be deleted",
		Description: "Description of task to be deleted",
		Status:      false,
		UserID:      userID,
	}

	mockTaskService.On("GetTaskByID", taskID).Return(existingTask, nil)
	mockTaskService.On("DeleteTask", taskID, userID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/tasks/"+taskID.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "Task deleted successfully", response["message"])

	mockTaskService.AssertExpectations(t)
}
