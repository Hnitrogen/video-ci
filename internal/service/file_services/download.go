package file_services

import (
	"awesomeProject1/internal/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func DownloadImg(c *gin.Context) {
	resp := app.Gin{C: c}
	filepath := c.Query("filepath")
	file, err := os.Open(filepath)

	if err != nil {
		resp.Response(http.StatusInternalServerError, "文件不存在", err.Error())
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "无法获取文件信息")
		return
	}

	// 设置响应头
	c.Header("Content-Disposition", "attachment; filename=fule.png")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	http.ServeContent(c.Writer, c.Request, filepath, fileInfo.ModTime(), file)
	resp.Response(http.StatusOK, "下载成功", "")
	return
}

func DownloadBigFile(c *gin.Context) {
	resp := app.Gin{C: c}
	filepath := c.Query("filepath")
	file, err := os.Open(filepath)

	if err != nil {
		resp.Response(http.StatusInternalServerError, "文件不存在", err.Error())
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot get file info"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileInfo.Name())
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", string(fileInfo.Size()))

	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)
	resp.Response(http.StatusOK, "下载成功", "")
	return
}
