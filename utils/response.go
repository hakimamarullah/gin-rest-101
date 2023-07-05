package utils

import "github.com/gin-gonic/gin"

func Response(message string, data interface{}) map[string]any {
	return gin.H{"message": message, "data": data}
}