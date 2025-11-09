package models

import (
	"log"

	"github.com/gin-gonic/gin"
)

type JSON_Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Result  any    `json:"result"`
	Error   any    `json:"error,omitempty"`
}

func ErrorResponse(c *gin.Context, stts int, msg string, err any) {
	log.Println(err)
	c.JSON(stts, JSON_Response{
		Success: false,
		Message: msg,
		Result:  nil,
	})
}
