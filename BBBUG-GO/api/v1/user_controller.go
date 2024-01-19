package api

import (
	"fmt"
	"log"
	"net/http"
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

type LoginRequest struct {
	UserAccount  string `json:"user_account" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type UserInfo struct {
	PassCount    int         `json:"pass_count"`
	PushCount    int         `json:"push_count"`
	UserAdmin    bool        `json:"user_admin"`
	UserId       int         `json:"user_id"`
	UserIcon     int         `json:"user_icon"`
	UserSex      int         `json:"user_sex"`
	UserName     string      `json:"user_name"`
	UserHead     string      `json:"user_head"`
	UserRemark   string      `json:"user_remark"`
	UserExtra    string      `json:"user_extra"`
	UserDevice   string      `json:"user_device"`
	UserTouchtip string      `json:"user_touchtip"`
	UserVip      string      `json:"user_vip"`
	UserGroup    int         `json:"user_group"`
	MyRoom       interface{} `json:"myRoom"`
	UserShutdown bool        `json:"user_shutdown"`
	UserSongdown bool        `json:"user_songdown"`
	UserGuest    bool        `json:"user_guest"`
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
		"code": 200,
		"msg":  "登录成功",
		"data": data,
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

	user, err := models.RegByLogin(req.UserAccount, req.UserPassword, ctx.ClientIP(), "")
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
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["access_token"] = token
	ctx.SetCookie("access_token", token, 3600, "/", "121.4.166.249", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": data,
	})
}

func GetMyInfo(ctx *gin.Context) {
	var bq BasicReq
	err := ctx.BindJSON(&bq)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}

	if bq.AccessToken == GuestToken {
		user := UserInfo{
			UserId:       -1,
			UserTouchtip: "Ghost",
			UserHead:     "new/images/nohead.jpg",
			UserAdmin:    false,
			MyRoom:       false,
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"data": user,
		})
		return
	}

	uid, err := middleware.GetUserId(bq.AccessToken)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  err,
		})
		return
	}

	u, err := models.GetUserInfoById(uid)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 404,
			"data": err,
		})
		return
	}
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
	}
	r, err := models.GetRoomByUser(u.UserId)
	if err == nil {
		ui.MyRoom = r
	} else {
		ui.MyRoom = false
	}
	if pcStr, err := config.GetCacheString("pass_song_card_user_" + strconv.Itoa(u.UserId)); err == nil {
		if pc, err := strconv.Atoi(pcStr); err == nil {
			ui.PassCount = pc
		}
	}
	if pcStr, err := config.GetCacheString("push_song_card_user_" + strconv.Itoa(u.UserId)); err == nil {
		if pc, err := strconv.Atoi(pcStr); err == nil {
			ui.PushCount = pc
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": ui,
	})
}

type onlineReq struct {
	BasicReq
	RoomId int `json:"room_id" binding:"required"`
}

func GetOnline(ctx *gin.Context) {
	online := onlineReq{}
	err := ctx.BindJSON(&online)
	if err != nil {
		log.Printf("获取在线信息失败, 错误:%s", err)
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	cs := websocket.GetClientsByRoom(online.RoomId)
	log.Print(cs)
	var users []UserInfo
	for _, c := range cs {
		var u UserInfo
		uid, err := strconv.Atoi(c.UserId)
		if err != nil {
			u.UserName = "来自" + util.GetAddrPath(c.Addr) + "的游客"
			users = append(users, u)
		} else {
			u, err = GetUserData(uid)
			if err != nil {
				continue
			}
			if _, err = config.GetCacheString(fmt.Sprintf("shutdown_room_%d_user_%d", online.RoomId, u.UserId)); err == nil {
				u.UserShutdown = true
			}
			if _, err = config.GetCacheString(fmt.Sprintf("songdown_room_%d_user_%d", online.RoomId, u.UserId)); err == nil {
				u.UserSongdown = true
			}
			if _, err = config.GetCacheString(fmt.Sprintf("guest_room_%d_user_%d", online.RoomId, u.UserId)); err == nil {
				u.UserGuest = true
			}
			users = append(users, u)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": users,
	})
}

func GetUserData(uid int) (UserInfo, error) {
	ui := UserInfo{}
	u, err := models.GetUserInfoById(uid)
	if err != nil {
		return ui, err
	}
	ui = UserInfo{
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
		UserGroup:    u.UserGroup,
		UserAdmin:    util.GetIsAdmin(u.UserGroup),
	}
	return ui, nil
}

type UpdateUserReq struct {
	BasicReq
	UserName     string `json:"user_name"`
	UserHead     string `json:"user_head"`
	UserRemark   string `json:"user_remark"`
	UserTouchtip string `json:"user_touchtip"`
	UserSex      string `json:"user_sex"`
	UserPassword string `json:"user_password"`
}

func UpdateUserInfo(ctx *gin.Context) {
	uur := UpdateUserReq{}
	err := ctx.BindJSON(&uur)
	if err != nil {
		log.Printf("参数不合法:%s", err)
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if uur.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	if uur.UserName != "" {
		if strings.Contains(uur.UserHead, ".gif") {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "用户名不能为空",
			})
			return
		}
	}
	if uur.UserHead != "" {
		if strings.Contains(uur.UserHead, ".gif") {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "头像不支持gif",
			})
			return
		}
	}
	successMsg := "资料更新成功"
	column := make(map[string]interface{})
	if uur.UserPassword != "" {
		if len(uur.UserPassword) < 6 || len(uur.UserPassword) > 16 {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "密码长度应该在6-16位之间",
			})
			return
		}
		salt := models.Salt()
		password := util.EncodeBySalt(uur.UserPassword, salt)
		successMsg = "密码更新成功"
		column["user_password"] = password
		column["user_salt"] = salt
	}
	uid, err := middleware.GetUserId(uur.AccessToken)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "无效的token",
		})
		return
	}
	column["user_name"] = uur.UserName
	column["user_head"] = uur.UserHead
	column["user_remark"] = uur.UserRemark
	column["user_sex"] = uur.UserSex
	column["user_touchtip"] = uur.UserTouchtip
	err = models.UpdateUser(uid, column)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "个人信息更新失败",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  successMsg,
	})
}
