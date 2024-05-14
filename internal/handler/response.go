package handler

import (
	"github.com/f1rdavsi/reporter/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	logger.Error.Println(message)
	c.AbortWithStatusJSON(status, errorResponse{
		Message: message,
	})
}
