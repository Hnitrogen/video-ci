package upload_services

import "github.com/gin-gonic/gin"

func WelcomePing(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}
