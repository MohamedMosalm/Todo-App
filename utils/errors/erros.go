package errors

import "net/http"

type AppError struct {
	Code    string
	Message string
	Details error
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

// Auth Errors
var ErrInvalidCredentials = &AppError{Code: "INVALID_CREDENTIALS", Message: "Invalid email or password", Status: http.StatusUnauthorized}
var ErrUserExists = &AppError{Code: "USER_EXISTS", Message: "User with this email already exists", Status: http.StatusConflict}
var ErrRegistrationFailed = &AppError{Code: "REGISTRATION_FAILED", Message: "Failed to register user", Status: http.StatusInternalServerError}
var ErrUserNotFound = &AppError{Code: "USER_NOT_FOUND", Message: "User not found", Status: http.StatusNotFound}
var ErrTokenGenerationFailed = &AppError{Code: "TOKEN_GENERATION_FAILED", Message: "Failed to generate access token", Status: http.StatusInternalServerError}

// Task Errors
var ErrInvalidTaskID = &AppError{Code: "INVALID_TASK_ID", Message: "Invalid task ID", Status: http.StatusBadRequest}
var ErrTaskNotFound = &AppError{Code: "TASK_NOT_FOUND", Message: "Task not found", Status: http.StatusNotFound}
var ErrCreateTaskFailed = &AppError{Code: "CREATE_FAILED", Message: "Failed to create task", Status: http.StatusInternalServerError}
var ErrFetchTasksFailed = &AppError{Code: "FETCH_ERROR", Message: "Failed to retrieve tasks", Status: http.StatusInternalServerError}
var ErrUpdateTaskFailed = &AppError{Code: "UPDATE_FAILED", Message: "Failed to update task", Status: http.StatusInternalServerError}
var ErrDeleteTaskFailed = &AppError{Code: "DELETE_FAILED", Message: "Failed to delete task", Status: http.StatusInternalServerError}

// General Errors
var ErrInvalidRequest = &AppError{Code: "INVALID_REQUEST", Message: "Invalid request body", Status: http.StatusBadRequest}
var ErrValidationError = &AppError{Code: "VALIDATION_ERROR", Message: "Validation failed", Status: http.StatusBadRequest}
var ErrUnauthorized = &AppError{Code: "UNAUTHORIZED", Message: "Unauthorized", Status: http.StatusUnauthorized}
var ErrInvalidUserID = &AppError{Code: "INVALID_USER_ID", Message: "Invalid user ID", Status: http.StatusBadRequest}
