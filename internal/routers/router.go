package routers

import (
	"awesomeProject1/internal/service/user_services"
	"github.com/gin-gonic/gin"
)
import "awesomeProject1/internal/service/upload_services"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", upload_services.WelcomePing)

	user := r.Group("/user")
	{
		user.POST("/login",user_services.)
		user.POST("/", user_services.CreateUser)
		//user.GET("/")
	}
	return r
}
