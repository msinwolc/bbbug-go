package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/msinwolc/api/v1"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.Index)
	r.POST("/user/sign_in", api.LoginUser)
	r.POST("/user/sign_up", api.RegistUser)

	return r
}
