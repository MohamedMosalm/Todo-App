package httputil

import (
	"log"

	"github.com/MohamedMosalm/Todo-App/utils/errors"
	"github.com/MohamedMosalm/Todo-App/utils/response"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err *errors.AppError) {
	resp := response.Response{
		Status: "error",
		Error: &response.ErrorInfo{
			Message: err.Message,
			Details: "",
		},
	}
	if err.Details != nil {
		log.Printf("Error details: %v", err.Details)
		resp.Error.Details = err.Details.Error()
	}
	c.JSON(err.Status, resp)
}

func SendSuccess(c *gin.Context, status int, message string, data interface{}) {
	resp := response.Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(status, resp)
}
