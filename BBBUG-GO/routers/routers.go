package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/msinwolc/api/v1"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	// r.Static("/new", "./new")

	r.GET("/", api.Index)

	r.POST("/user/login", api.LoginUser)
	// r.POST("/user/sign_up", api.RegistUser)
	r.POST("/user/getmyinfo", api.GetMyInfo)
	r.POST("/user/online", api.GetOnline)

	r.POST("/room/getRoomInfo", api.GetRoomInfo)
	r.POST("/room/getWebsocketUrl", api.GetWebsocketUrl)

	r.POST("/message/getMessageList", api.GetMessageList)
	r.POST("/message/send", api.Send)

	r.POST("/song/search", api.SearchSong)
	r.POST("/song/addSong", api.AddSong)
	r.POST("/song/songList", api.SongList)
	r.POST("/song/push", api.PushSong)
	r.POST("/song/playSong", api.PlaySong)
	r.POST("/song/playurl", api.PlayUrl)
	r.GET("/song/playurl", api.PlayUrlGet)
	r.POST("/song/getLrc", api.GetLrc)
	r.POST("/song/remove", api.RemoveSong)

	r.POST("/system/time", api.GetTime)

	r.GET("/socket", api.ConnWebsocket)

	return r
}
