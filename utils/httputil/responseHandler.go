package httputil

import (
	"github.com/MohamedMosalm/To-Do-List/utils/errors"
	"github.com/MohamedMosalm/To-Do-List/utils/response"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleError(c *gin.Context, err *errors.AppError) {
	resp := response.Response{
		Status: "error",
		Error: &response.ErrorInfo{
			Code:    err.Code,
			Message: err.Message,
			Details: "",
		},
	}
	if err.Details != nil {
		log.Printf("Error details: %v", err.Details)
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
