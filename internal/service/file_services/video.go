package file_services

import (
	"awesomeProject1/internal/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
)

func PlayList(c *gin.Context) {
	resp := app.Gin{C: c}

	inputFile := "storage/chunk/1720339336744/dist.mp4"
	outputPrefix := "storage/chunk/1720339336744/output/"

	absInputFile, err := filepath.Abs(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// 输出文件前缀（相对或绝对路径）
	absOutputPrefix, err := filepath.Abs(outputPrefix)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(absInputFile, absOutputPrefix)

	//segmentSize := 3
	// TODO limit cpu
	cmd := exec.Command("ffmpeg", "-i", absInputFile, "-c:v", "copy",
		"-c:a", "copy", "-f", "segment", "-segment_time", strconv.Itoa(3), "-segment_list",
		absOutputPrefix+"/playlist.m3u8", "-segment_format", "mpegts", "-threads", "4", outputPrefix+"/output%03d.ts")

	err = cmd.Start()

	if err != nil {
		resp.Response(http.StatusInternalServerError, err.Error(), "")
	}

	err = cmd.Wait()
	if err != nil {
		resp.Response(http.StatusInternalServerError, err.Error(), "")
	}
	resp.Response(http.StatusOK, "fule", "")
	return

}
