package utils

import (
	"fmt"

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
	responseBody := StdResponse{
		Status:  "success",
		Message: msg,
		Data:    data,
	}
	fmt.Println(responseBody)
	c.JSON(statusCode, responseBody)
}

func FailureResponse(c *gin.Context, statusCode int, msg string, err interface{}) {
	if err != nil {
		err = ErrorResponse{120, err}
	}
	responseBody := StdResponse{
		Status:  "failure",
		Message: msg,
		Error:   err,
	}
	fmt.Println(responseBody)
	c.JSON(statusCode, responseBody)
}
