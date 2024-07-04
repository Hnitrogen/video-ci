package routers

import (
	"awesomeProject1/internal/service/file_services"
	"awesomeProject1/internal/service/user_services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
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
		upload.POST("/", file_services.UploadImg)
		upload.POST("/chunk", file_services.UploadFileChunk)
		upload.GET("/merge", file_services.MergeFileChunk)
	}

	// DownloadRouters
	download := r.Group("/download")
	{
		download.GET("/img", file_services.DownloadImg)
	}
	return r
}
