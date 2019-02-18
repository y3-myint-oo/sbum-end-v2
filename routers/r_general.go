package routers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var pool = "ABCDEFGHIJKLMNOPQRSTUVWHYZ0123456789"

// Generate a random string of A-Z chars with len = l
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

//////////////////////////////////////////////////////////////

// Serail Number Generator
func serialNumber(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": randomString(15),
	})
}
