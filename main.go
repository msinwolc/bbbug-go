package main

import (
	"log"

	"github.com/msinwolc/routers"

	"github.com/msinwolc/models"
)

func main() {
	// r := gin.Default()
	// r.LoadHTMLGlob("templates/**/*")
	// r.GET("/post/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "post.html", gin.H{
	// 		"title": "This is a post",
	// 	})
	// })

	// r.GET("/post/new", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "form.html", gin.H{})
	// })
	// r.POST("/post/create", func(ctx *gin.Context) {
	// 	types := ctx.DefaultPostForm("type", "post")
	// 	username := ctx.PostForm("username")
	// 	password := ctx.PostForm("userpassword")
	// 	ctx.String(http.StatusOK, fmt.Sprintf("type: %s, username: %s, password: %s", types, username, password))
	// })

	// r.GET("/post/upload", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "upload.html", gin.H{})
	// })
	// r.POST("/post/upload", func(ctx *gin.Context) {
	// 	file, err := ctx.FormFile("file")
	// 	if err != nil {
	// 		ctx.String(http.StatusBadRequest, "上传文件错误")
	// 	}
	// 	ctx.SaveUploadedFile(file, file.Filename)
	// 	ctx.String(http.StatusOK, file.Filename)
	// })
	models.GetDB()

	models.OpenDB()

	defer models.CloseDB()
	r := routers.Routers()
	// fmt.Println(r)
	if err := r.Run(); err != nil {
		log.Fatalf("startup server failed: %v", err)
	}
}
