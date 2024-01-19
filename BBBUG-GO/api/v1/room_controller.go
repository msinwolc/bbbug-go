package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msinwolc/config"
	"github.com/msinwolc/middleware"
	"github.com/msinwolc/models"
	"github.com/msinwolc/util"
)

type RoomInfoReq struct {
	BasicReq
	RoomId       string `json:"room_id" binding:"required"`
	RoomPassword string `json:"room_password"`
}

func GetRoomInfo(ctx *gin.Context) {
	rq := RoomInfoReq{}
	err := ctx.BindJSON(&rq)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	room_id, err := strconv.Atoi(rq.RoomId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	room, err := models.GetRoomById(room_id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "未查询到房间",
		})
		return
	}
	if room.RoomStatus == 1 {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  room.RoomReason,
		})
		return
	}
	if rq.AccessToken == GuestToken {
		if room.RoomPublic == 1 {
			ctx.JSON(200, gin.H{
				"code": 401,
				"msg":  "该房间为私密房间，游客无法访问",
			})
			return
		}
		room.RoomPassword = ""
		u, err := models.GetUserInfoById(room.RoomUser)
		if err != nil {
			room.Admin = false
		} else {
			ui := UserInfo{
				UserId:       u.UserId,
				UserTouchtip: u.UserTouchtip,
				UserSex:      u.UserSex,
				UserRemark:   u.UserRemark,
				UserHead:     u.UserHead,
				UserName:     u.UserName,
				UserDevice:   u.UserDevice,
				UserExtra:    u.UserExtra,
				UserIcon:     u.UserIcon,
				UserVip:      u.UserVip,
				UserAdmin:    true,
				MyRoom:       false,
			}
			room.Admin = ui
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"data": room,
		})
		return
	}
	uid, err := middleware.GetUserId(rq.AccessToken)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  err,
		})
		return
	}
	if room.RoomPublic == 1 && room.RoomUser != uid {
		cacheRoomPwdName := fmt.Sprintf("password_room_%d_password_%d", room.RoomId, uid)
		savedPassword, err := config.GetCacheString(cacheRoomPwdName)
		if err != nil {
			savedPassword = ""
		}
		if savedPassword != room.RoomPassword && rq.RoomPassword != room.RoomPassword {
			ctx.JSON(200, gin.H{
				"code": 401,
				"msg":  "密码错误",
			})
			return
		}
		err = config.SetCacheString(cacheRoomPwdName, room.RoomPassword, 86400*time.Second)
		if err != nil {
			log.Printf("缓存房间密码失败, 错误:%s", err)
		}
	}
	room.RoomPassword = ""
	u, err := models.GetUserInfoById(room.RoomUser)
	if err != nil {
		room.Admin = false
	} else {
		ui := UserInfo{
			UserId:       u.UserId,
			UserTouchtip: u.UserTouchtip,
			UserSex:      u.UserSex,
			UserRemark:   u.UserRemark,
			UserHead:     u.UserHead,
			UserName:     u.UserName,
			UserDevice:   u.UserDevice,
			UserExtra:    u.UserExtra,
			UserIcon:     u.UserIcon,
			UserVip:      u.UserVip,
			UserAdmin:    true,
			MyRoom:       false,
		}
		room.Admin = ui
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": room,
	})
}

type WebsocketUrlReq struct {
	BasicReq
	Channel      int    `json:"channel" binding:"required"`
	RoomPassword string `json:"room_password"`
}

func GetWebsocketUrl(ctx *gin.Context) {
	wsurl := WebsocketUrlReq{}
	err := ctx.BindJSON(&wsurl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数不合法",
		})
		return
	}
	_, err = models.GetRoomById(wsurl.Channel)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "未找到房间",
		})
		return
	}
	ip := ctx.Request.RemoteAddr

	addr, err := config.GetCacheString("ip_addr_" + ip)
	if err != nil || addr == "" {
		addr = util.GetAddrPath(ip)
		err = config.SetCacheString("ip_addr_"+ip, addr, 3600*time.Second)
		if err != nil {
			log.Printf("缓存ip地址失败,错误:%s", err)
		}
	}
	lastSend, err := config.GetCacheString("channel_" + strconv.Itoa(wsurl.Channel) + "_user_" + ip)
	if err != nil || lastSend == "" {
		content := "欢迎"
		if addr != "" {
			content += "来自" + content + "的"
		}
		if wsurl.Plat != "" {
			content += wsurl.Plat + "用户"
		} else {
			content += "临时用户"
		}
		msg := make(map[string]interface{})
		msg["type"] = "join"
		msg["name"] = "临时用户"
		msg["where"] = addr
		msg["plat"] = wsurl.Plat
		msg["content"] = content
		err = config.SetCacheString("channel_"+strconv.Itoa(wsurl.Channel)+"_user_"+ip, content, 3600*time.Second)
		if err != nil {
			log.Printf("缓存消息失败, 错误:%s\n", err)
		}
	}
	data := make(map[string]interface{})
	data["account"] = ip
	data["channel"] = wsurl.Channel
	data["ticket"] = util.GetWebsocketTicket(ip, wsurl.Channel)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

type listReq struct {
	BasicReq
	Rid int `json:"room_id" binding:"required"`
}

func SongList(ctx *gin.Context) {
	var lr listReq
	err := ctx.BindJSON(&lr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	var lss = GetSongListFromCache(lr.Rid)
	if len(lss) == 0 {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": []ListSong{},
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": lss,
	})
}
