package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/msinwolc/controllers"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.Index)

	return r
}
