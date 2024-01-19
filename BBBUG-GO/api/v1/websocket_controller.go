package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/msinwolc/util"
	socket "github.com/msinwolc/websocket"
)

func ConnWebsocket(ctx *gin.Context) {
	channel := ctx.Query("channel")
	rid, err := strconv.Atoi(channel)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if util.GetWebsocketTicket(ctx.Query("account"), rid) != ctx.Query("ticket") {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "客户端登录失败",
		})
		return
	}
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("连接出错:%s", err)
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "连接出错",
		})
		return
	}
	socket.NewClient(conn, conn.RemoteAddr().String(), ctx.Query("account"), ctx.Query("ticket"), rid)
}
