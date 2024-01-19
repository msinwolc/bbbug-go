package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msinwolc/config"
	"github.com/msinwolc/middleware"
	"github.com/msinwolc/models"
	"github.com/msinwolc/util"
	"github.com/msinwolc/websocket"
)

type msgListReq struct {
	BasicReq
	RoomId       int    `json:"room_id" binding:"required"`
	MessageWhere string `json:"message_where"`
	Page         int    `json:"page"`
	PerPage      int    `json:"per_page"`
}

func GetMessageList(ctx *gin.Context) {
	mlr := msgListReq{}
	err := ctx.BindJSON(&mlr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数不合法",
		})
		return
	}
	if mlr.MessageWhere == "" {
		mlr.MessageWhere = "channel"
	}
	r, err := models.GetRoomById(mlr.RoomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "房间获取失败",
		})
		return
	}
	if mlr.Page == 0 {
		mlr.Page = 1
	}
	if mlr.PerPage == 0 {
		mlr.PerPage = 100
	}
	if mlr.Page > 100 {
		mlr.PerPage = 100
	}
	roomMessageListCacheName := fmt.Sprintf("room_message_list_%d", mlr.RoomId)
	// var mlist []models.Message
	_, err = config.GetCacheString(roomMessageListCacheName)
	if err != nil {
		mlsMap := make(map[string]interface{})
		mlsMap["message_to"] = strconv.Itoa(r.RoomId)
		mlsMap["message_where"] = mlr.MessageWhere
		mlsMap["message_status"] = 0
		mls := models.GetMessageListByMap(mlsMap, mlr.Page, mlr.PerPage)
		mlsByte, err := json.Marshal(mls)
		if err != nil {
			log.Println("json.Marshal(mls)失败")
		}
		err = config.SetCacheString(roomMessageListCacheName, string(mlsByte), 10*time.Second)
		if err != nil {
			log.Printf("SetCacheString失败, 错误:%s", err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": mls,
		})
		return
	}

}

type sendReq struct {
	BasicReq
	Where    string      `json:"where" binding:"required"`
	To       int         `json:"to" binding:"required"`
	Type     string      `json:"type" binding:"required"`
	At       interface{} `json:"at"`
	Msg      string      `json:"msg" binding:"required"`
	Resource string      `json:"resource"`
}

func Send(ctx *gin.Context) {
	var sr sendReq
	err := ctx.BindJSON(&sr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if sr.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(sr.AccessToken)
	if err != nil {
		log.Printf("获取用户信息失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户信息获取失败",
		})
		return
	}
	u, err := GetUserData(uid)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户信息获取失败",
		})
		return
	}
	r, err := models.GetRoomById(sr.To)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "房间信息获取失败",
		})
		return
	}
	savedPassword := ""
	if cachePassword, err := config.GetCacheString(fmt.Sprintf("password_room_%d_password_%d", r.RoomId, uid)); err == nil {
		savedPassword = cachePassword
	}
	if r.RoomPublic == 1 && uid != r.RoomUser && savedPassword != r.RoomPassword {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "密码错误",
		})
		return
	}
	if r.RoomSendmsg == 1 && uid != r.RoomUser {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "全员禁言中",
		})
		return
	}
	msg := strings.Replace(sr.Msg, " ", "", -1)
	if len(msg) == 0 {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "消息内容不能为空",
		})
		return
	}
	isban := false
	if _, err = config.GetCacheString(fmt.Sprintf("shutdown_room_%d_user_%d", sr.To, uid)); err == nil {
		isban = true
	}
	switch sr.Where {
	case "channel":
		if isban && uid > 1 && util.GetIsAdmin(u.UserGroup) {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "你被禁止发言",
			})
			return
		}
	default:
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "@#%^$&%",
		})
		return
	}
	if util.GetIsAdmin(u.UserGroup) {
		if strings.Contains(sr.Msg, "@all") {
			sendAllMap := make(map[string]interface{})
			sendAllMap["type"] = "all"
			sendAllMap["content"] = strings.Replace(sr.Msg, "@all", "", -1)
			websocket.SendMsgToAll(sendAllMap)
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  "已发送",
			})
			return
		}
	} else {
		ip := ctx.Request.RemoteAddr
		if _, err = config.GetCacheString("black_ip_" + ip); err == nil {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "你的IP被ban，无法发送消息",
			})
			return
		}
		switch sr.Type {
		case "text":
			if uid != r.RoomUser {
				if len(sr.Msg) > 250 {
					ctx.JSON(200, gin.H{
						"code": 500,
						"msg":  "超过最大字数限制(250)",
					})
					return
				}
				if oldMsg, err := config.GetCacheString(fmt.Sprintf("last_%d", uid)); err == nil {
					if oldMsg == sr.Msg {
						ctx.JSON(200, gin.H{
							"code": 500,
							"msg":  "你是复读机吗",
						})
						return
					}
				}
			}
		}
	}
	msgUrl, _ := url.QueryUnescape(sr.Msg)
	switch sr.Type {
	case "text":
		if strings.Contains(msgUrl, "121.4.166.249") {
			regUrl := regexp.MustCompile(`/(\d+)`)
			rids := regUrl.FindAllStringSubmatch(msgUrl, -1)
			rid := 0
			if len(rids) > 0 {
				ridStr := rids[len(rids)-1][1]
				rid, _ = strconv.Atoi(ridStr)
			}
			if jumpRoom, err := models.GetRoomById(rid); err == nil {
				_, err := config.GetCacheString(fmt.Sprintf("chat_message_jump_%d", uid))
				if err == nil && !util.GetIsAdmin(u.UserGroup) {
					ctx.JSON(200, gin.H{
						"code": 500,
						"msg":  "发送机票过于频繁",
					})
					return
				} else {
					err = config.SetCacheString(fmt.Sprintf("chat_message_jump_%d", uid), "true", 60*time.Second)
					if err != nil {
						log.Printf("缓存机票信息失败:%s", err)
					}
					if jumpRoom.RoomPassword == "" {
						jumpRoom.RoomPassword = "true"
					} else {
						jumpRoom.RoomPassword = "false"
					}
					m := models.Message{
						MessageUser:    uid,
						MessageType:    "text",
						MessageContent: sr.Msg,
						MessageStatus:  1,
						MessageWhere:   sr.Where,
						MessageTo:      strconv.Itoa(r.RoomId),
					}
					m, err = models.CreateMessage(m)
					if err != nil {
						ctx.JSON(200, gin.H{
							"code": 500,
							"msg":  "发送失败",
						})
						return
					}
					sendMsg := make(map[string]interface{})
					sendMsg["type"] = "jump"
					sendMsg["jump"] = jumpRoom
					sendMsg["message_id"] = m.MessageId
					sendMsg["message_time"] = time.Now().Unix()
					sendMsg["user"] = u
					websocket.SendMessageToRoom(r.RoomId, sendMsg)
					contentMap := make(map[string]interface{})
					contentMap["type"] = "text"
					contentMap["where"] = sr.Where
					contentMap["at"] = sr.At
					contentMap["message_id"] = m.MessageId
					contentMap["message_time"] = time.Now().Unix()
					contentMap["content"] = sr.Msg
					contentMap["resource"] = sr.Resource
					contentMap["user"] = u
					updateMessage := make(map[string]interface{})
					updateMessage["message_type"] = "text"
					if contentByte, err := json.Marshal(contentMap); err == nil {
						updateMessage["message_content"] = string(contentByte)
					}
					updateMessage["message_status"] = 0
					err = models.UpdateMessageById(m.MessageId, updateMessage)
					if err != nil {
						log.Printf("更新message数据时失败:%s", err)
					}
					ctx.JSON(200, gin.H{
						"code": 200,
						"msg":  "发送机票成功",
					})
					return
				}
			}
		}
		m := models.Message{
			MessageUser:    uid,
			MessageType:    "text",
			MessageContent: sr.Msg,
			MessageStatus:  1,
			MessageWhere:   sr.Where,
			MessageTo:      strconv.Itoa(r.RoomId),
		}
		m, err = models.CreateMessage(m)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "发送消息失败",
			})
			return
		}
		sendMsg := make(map[string]interface{})
		sendMsg["type"] = "text"
		sendMsg["where"] = sr.Where
		sendMsg["resource"] = sr.Resource
		sendMsg["at"] = sr.At
		sendMsg["content"] = sr.Msg
		sendMsg["message_id"] = m.MessageId
		sendMsg["message_time"] = time.Now().Unix()
		sendMsg["user"] = u
		websocket.SendMessageToRoom(r.RoomId, sendMsg)
		contentMap := make(map[string]interface{})
		contentMap["type"] = "text"
		contentMap["where"] = sr.Where
		contentMap["at"] = sr.At
		contentMap["message_id"] = m.MessageId
		contentMap["message_time"] = time.Now().Unix()
		contentMap["content"] = sr.Msg
		contentMap["resource"] = sr.Resource
		contentMap["user"] = u
		updateMessage := make(map[string]interface{})
		updateMessage["message_type"] = "text"
		if contentByte, err := json.Marshal(contentMap); err == nil {
			updateMessage["message_content"] = string(contentByte)
		}
		updateMessage["message_status"] = 0
		err = models.UpdateMessageById(m.MessageId, updateMessage)
		if err != nil {
			log.Printf("更新message数据时失败:%s", err)
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "发送消息成功",
		})
		return
	}
}
