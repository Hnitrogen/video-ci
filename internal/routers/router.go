package routers

import (
	"awesomeProject1/internal/service/file_services"
	"awesomeProject1/internal/service/user_services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("conf/*")

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
		upload.GET("/clean", file_services.CleanChunk)
	}

	// DownloadRouters
	download := r.Group("/download")
	{
		download.GET("/img", file_services.DownloadImg)
		download.GET("/big", file_services.DownloadBigFile)
	}

	// templates
	templates := r.Group("/templates")
	{
		templates.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "big_file_upload.html", gin.H{})
		})
	}

	return r
}
