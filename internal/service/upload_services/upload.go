package upload_services

import (
	"awesomeProject1/internal/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	uploads_map = make(map[string]*os.File)
	uploadsLock sync.Mutex
)

func UploadImg(c *gin.Context) {
	resp := app.Gin{C: c}
	_, image, err := c.Request.FormFile("image")
	if err != nil {
		resp.Response(http.StatusInternalServerError, "服务异常", "")
		return
	}

	if image == nil {
		resp.Response(http.StatusInternalServerError, "服务异常", "")
		return
	}

	filename := image.Filename
	if c.SaveUploadedFile(image, "storage/"+filename) != nil {
		resp.Response(http.StatusInternalServerError, "图片保存失败", "")
		return
	}
	resp.Response(http.StatusOK, "文件保存成功", "")
	return
}

func UploadFileChunk(c *gin.Context) {
	resp := app.Gin{C: c}
	file, _, err := c.Request.FormFile("chunk")

	if err != nil {
		resp.Response(http.StatusInternalServerError, "分片上传失败", "")
		return
	}
	defer file.Close()

	idx := c.Request.FormValue("idx")
	md5 := c.Request.FormValue("md5")
	chunkDir := "storage/chunk/" + md5
	err = os.MkdirAll(chunkDir, os.ModePerm)

	if err != nil {
		resp.Response(http.StatusInternalServerError, "临时文件夹创建失败", "")
		return
	}

	chunkPath := filepath.Join(chunkDir, idx)
	//chunkPath := chunkDir
	chunkFile, err := os.Create(chunkPath)

	if err != nil {
		resp.Response(http.StatusInternalServerError, "分片创建失败", err)
		return
	}
	defer chunkFile.Close()

	_, err = io.Copy(chunkFile, file)
	if err != nil {
		resp.Response(http.StatusInternalServerError, "分片文件写入失败", "")
		return
	}

	//TODO add lock
	//uploads_map[hash] = chunkFile

	resp.Response(http.StatusOK, "分片保存成功", "")
}

func MergeFileChunk(c *gin.Context) {
	resp := app.Gin{C: c}
	chunkPath := c.Query("chunkPath")
	distPath := chunkPath + "/dist"
	filelist, err := os.ReadDir(chunkPath)
	// if dist.mp4 in filelist , file merge cannot stop
	distFile, err := os.Create(distPath + ".mp4")

	if err != nil {
		resp.Response(http.StatusInternalServerError, "合并文件初始化失败", err)
		return
	}
	defer distFile.Close()

	//merge chunks
	for idx, file := range filelist {
		fmt.Printf("copy idx = %v --- filename = %v\n", idx, file.Name())

		chunk, _ := os.Open(filepath.Join(chunkPath, file.Name()))
		io.Copy(distFile, chunk)
		chunk.Close()
	}

	if err != nil {
		resp.Response(http.StatusInternalServerError, "合并切片异常", err)
		return
	}

	resp.Response(http.StatusOK, "合并成功", "")
	return
}
