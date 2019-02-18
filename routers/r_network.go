package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func networkStatus(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Your are online",
	})
}
