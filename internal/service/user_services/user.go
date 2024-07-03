package user_services

import (
	"awesomeProject1/internal/app"
	models "awesomeProject1/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

type UserDTO struct {
	username string
	password string
}

func Login(context *gin.Context) {
	resp := app.Gin{
		C: context,
	}

	username := context.PostForm("username")
	password := context.PostForm("password")

	if models.CheckUser(username, password) {
		resp.Response(http.StatusOK, 200, "登录成功", "")
	} else {
		resp.Response(http.StatusUnauthorized, 401, "账号或密码错误", "")
	}
}

func CreateUser(context *gin.Context) {
	resp := app.Gin{
		C: context,
	}
	username := context.PostForm("username")
	password := context.PostForm("password")

	if username == "" || password == "" {
		resp.Response(http.StatusInternalServerError, 500, "参数校验失败", "")
		log.Printf("参数为空: username = %s , password = %s", username, password)
		return
	}

	models.AddUser(username, password)
	resp.Response(http.StatusCreated, 200, "创建成功", "")
	return
}
