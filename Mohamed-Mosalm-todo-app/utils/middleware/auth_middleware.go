package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/MohamedMosalm/Todo-App/utils/errors"
	"github.com/MohamedMosalm/Todo-App/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			appErr := errors.ErrUnauthorized
			appErr.Details = err
			httputil.HandleError(c, appErr)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > exp {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			httputil.HandleError(c, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			appErr := errors.ErrUnauthorized
			appErr.Details = err
			httputil.HandleError(c, appErr)
			c.Abort()
			return
		}

		c.Set("user_id", userID.String())
		c.Next()
	}
}
