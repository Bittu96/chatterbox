package utils

import (
	"github.com/gin-gonic/gin"
)

type StdResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, msg string, data interface{}) {
	c.JSON(statusCode, StdResponse{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}

func FailureResponse(c *gin.Context, statusCode int, msg string, err interface{}) {
	if err != nil {
		err = ErrorResponse{120, err}
	}
	c.JSON(statusCode, StdResponse{
		Status:  "failure",
		Message: msg,
		Error:   err,
	})
}
