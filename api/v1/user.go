package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msinwolc/config"
	"github.com/msinwolc/middleware"
	"github.com/msinwolc/models"
)

type LoginRequest struct {
	UserAccount  string `json:"user_account" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
	UserDevice   string `json:"user_device" binding:"required"`
}

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Welcome",
	})
}

func LoginUser(ctx *gin.Context) {
	req := LoginRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "非法参数",
		})
		return
	}

	user, err := models.Login(req.UserAccount, req.UserPassword)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf(err.Error()),
		})
		return
	}
	token, err := middleware.CreateToken(req.UserAccount, "test", user.UserId, 24*7*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	err = config.SetToken(user.UserName, token)
	if err != nil {
		fmt.Println(err)
	}
	err = models.CreateAccess(user.UserId, "test", token, ctx.ClientIP())
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]interface{})
	data["access_token"] = token
	ctx.SetCookie("access_token", token, 3600, "/", "1.14.193.52", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": user,
	})
}

func RegistUser(ctx *gin.Context) {
	req := LoginRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "非法参数",
		})
		return
	}

	user, err := models.RegByLogin(req.UserAccount, req.UserPassword, ctx.ClientIP(), req.UserDevice)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf(err.Error()),
		})
		return
	}
	token, err := middleware.CreateToken(req.UserAccount, "test", user.UserId, 24*7*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	err = config.SetToken(user.UserName, token)
	if err != nil {
		fmt.Println(err)
	}
	err = models.CreateAccess(user.UserId, "test", token, ctx.ClientIP())
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]interface{})
	data["access_token"] = token
	ctx.SetCookie("access_token", token, 3600, "/", "1.14.193.52", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": user,
	})
}
