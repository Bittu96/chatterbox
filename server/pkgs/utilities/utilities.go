package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		fmt.Println("envs", os.Environ())
		fmt.Println("loaded", len(os.Environ()), "envs")
	}
}

func GetEnv(key string) (value string) {
	if value = os.Getenv(key); value == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file", err)
		} else {
			if value = os.Getenv(key); value == "" {
				log.Fatalf("env key %v not found", key)
			} else {
				fmt.Println("envs", os.Environ())
				fmt.Println("loaded", len(os.Environ()), "envs")
			}
		}
	}
	return
}

func ToInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func ToFloat(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err != nil {
		panic(err)
	} else {
		return f
	}
}

func ToBool(s string) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		panic(err)
	} else {
		return b
	}
}
