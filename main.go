package main

import (
	"awesomeProject1/internal/routers"
	"awesomeProject1/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	setting.Setup() // 初始化ini
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Printf("%v", endPoint)
	server.ListenAndServe()

}
