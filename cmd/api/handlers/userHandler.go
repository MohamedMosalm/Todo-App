package handlers

import (
	"net/http"

	"github.com/MohamedMosalm/To-Do-List/dtos"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/MohamedMosalm/To-Do-List/services"
	"github.com/MohamedMosalm/To-Do-List/utils/auth"
	"github.com/MohamedMosalm/To-Do-List/utils/errors"
	"github.com/MohamedMosalm/To-Do-List/utils/httputil"
	"github.com/MohamedMosalm/To-Do-List/utils/validator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService     services.UserService
	authValidator   *validator.AuthValidator
	jwtService      *auth.JWTService
	passwordService *auth.PasswordService
}

func NewAuthHandler(userService services.UserService) (*AuthHandler, error) {
	jwtService, err := auth.NewJWTService()
	if err != nil {
		return nil, err
	}

	return &AuthHandler{
		userService:     userService,
		authValidator:   validator.NewAuthValidator(),
		jwtService:      jwtService,
		passwordService: auth.NewPasswordService(),
	}, nil
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerDTO dtos.RegisterDTO

	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		appErr := errors.ErrInvalidRequest
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	if err := h.authValidator.ValidateRegisterDTO(&registerDTO); err != nil {
		appErr := errors.ErrValidationError
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	existingUser, _ := h.userService.FindUserByEmail(registerDTO.Email)
	if existingUser != nil {
		httputil.HandleError(c, errors.ErrUserExists)
		return
	}

	hashedPassword, err := h.passwordService.HashPassword(registerDTO.Password)
	if err != nil {
		appErr := errors.ErrRegistrationFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	user := models.User{
		FirstName: registerDTO.FirstName,
		LastName:  registerDTO.LastName,
		Email:     registerDTO.Email,
		Phone:     registerDTO.Phone,
		Password:  hashedPassword,
	}

	if err := h.userService.CreateUser(&user); err != nil {
		appErr := errors.ErrRegistrationFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	httputil.SendSuccess(c, http.StatusCreated, "User registered successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginDTO dtos.LoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		appErr := errors.ErrInvalidRequest
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	if err := h.authValidator.ValidateLoginDTO(&loginDTO); err != nil {
		appErr := errors.ErrValidationError
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	user, err := h.userService.FindUserByEmail(loginDTO.Email)
	if err != nil {
		httputil.HandleError(c, errors.ErrUserNotFound)
		return
	}

	if err := h.passwordService.ComparePasswords(user.Password, loginDTO.Password); err != nil {
		httputil.HandleError(c, errors.ErrInvalidCredentials)
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		appErr := errors.ErrTokenGenerationFailed
		appErr.Details = err
		httputil.HandleError(c, appErr)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, "Login successful", gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	})
}
