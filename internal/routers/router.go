package routers

import (
	"awesomeProject1/internal/service/upload_services"
	"awesomeProject1/internal/service/user_services"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// UserRouters
	user := r.Group("/user")
	{
		user.POST("/login", user_services.Login)
		user.POST("/", user_services.CreateUser)
		//user.GET("/")
	}

	// UploadRouters
	upload := r.Group("/upload")
	{
		upload.POST("/", upload_services.UploadImg)
	}
	return r
}
