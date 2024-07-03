package upload_services

import (
	"awesomeProject1/internal/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImg(c *gin.Context) {
	resp := app.Gin{C: c}
	_, image, err := c.Request.FormFile("image")
	if err != nil {
		resp.Response(http.StatusBadRequest, "服务异常", "")
		return
	}

	if image == nil {
		resp.Response(http.StatusBadRequest, "服务异常", "")
		return
	}

	c.SaveUploadedFile(image, "storage/fule.webp")
	resp.Response(http.StatusOK, "文件保存成功", "")
	return
}
