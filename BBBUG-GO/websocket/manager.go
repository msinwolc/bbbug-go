package websocket

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/msinwolc/config"
	"github.com/msinwolc/models"
	"github.com/msinwolc/util"
)

type ClientManager struct {
	Clients     map[*Client]struct{}
	ClientsLock *sync.RWMutex
	Register    chan *Client
	Unregister  chan *Client
}

type SocketMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Where   string `json:"where"`
	At      string `json:"at"`
	Name    string `json:"name"`
	Plat    string `json:"plat"`
}

var SCM *ClientManager

func (client *Client) UnConnectSocket() {
	err := client.Socket.Close()
	if err != nil {
		log.Printf("%v(%s)房号:%d断开连接时出错", client.UserId, client.Addr, client.RoomId)
	}
}

func InitSocketManager() {
	SCM = &ClientManager{
		Clients:     make(map[*Client]struct{}),
		ClientsLock: new(sync.RWMutex),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
	}
	go startSocket()
}

func startSocket() {
	for {
		select {
		case client := <-SCM.Register:
			client.EventRegister()
			log.Printf("现有总连接数:%d", len(SCM.Clients))
		case client := <-SCM.Unregister:
			client.EventUnRegister()
			log.Printf("现有总连接数:%d", len(SCM.Clients))
		}
	}
}

func (client *Client) EventRegister() {
	oldClients := GetClientsByTicket(client.Ticket)
	for _, c := range oldClients {
		c.UnConnectSocket()
		c.DelClientFromManager()
	}
	client.SetClientToManager()

	RefreshOnlineClientsByRoom(client.RoomId)
}

func (client *Client) EventUnRegister() {
	client.DelClientFromManager()
	RefreshOnlineClientsByRoom(client.RoomId)
}

func GetClientsByTicket(ticket string) []*Client {
	cs := []*Client{}
	SCM.ClientsLock.Lock()
	defer SCM.ClientsLock.Unlock()
	for c, _ := range SCM.Clients {
		if c.Ticket == ticket {
			cs = append(cs, c)
		}
	}
	return cs
}

func GetClientsByRoom(roomId int) []*Client {
	cs := []*Client{}
	SCM.ClientsLock.Lock()
	defer SCM.ClientsLock.Unlock()
	for c, _ := range SCM.Clients {
		if c.RoomId == roomId {
			cs = append(cs, c)
		}
	}
	return cs
}

func (client *Client) DelClientFromManager() {
	SCM.ClientsLock.Lock()
	defer SCM.ClientsLock.Unlock()
	if _, ok := SCM.Clients[client]; ok {
		delete(SCM.Clients, client)
	}
}

func (client *Client) SetClientToManager() {
	SCM.ClientsLock.Lock()
	defer SCM.ClientsLock.Unlock()
	SCM.Clients[client] = struct{}{}
}

type OnlineUser struct {
	UserId       int    `json:"user_id"`
	UserIcon     int    `json:"user_icon"`
	UserAdmin    bool   `json:"user_admin"`
	UserShutdown bool   `json:"user_shutdown"`
	UserSongdown bool   `json:"user_songdown"`
	UserGuest    bool   `json:"user_guest"`
	UserSex      int    `json:"user_sex"`
	UserName     string `json:"user_name"`
	UserHead     string `json:"user_head"`
	UserRemark   string `json:"user_remark"`
	UserExtra    string `json:"user_extra"`
	UserDevice   string `json:"user_device"`
	UserTouchtip string `json:"user_touchtip"`
	UserVip      string `json:"user_vip"`
	UserGroup    int    `json:"user_group"`
}

func RefreshOnlineClientsByRoom(roomId int) {
	cs := GetClientsByRoom(roomId)
	var users []OnlineUser
	online := len(cs)
	for _, c := range cs {
		uid, err := strconv.Atoi(c.UserId)
		if err != nil {
			name := util.GetAddrPath(c.Addr)
			users = append(users, OnlineUser{
				UserId:   888,
				UserName: name + "的游客",
			})
			online -= 1
		} else {
			user_data, err := GetOnlineUserData(uid)
			if err != nil {
				log.Printf("获取用户信息失败:%s", err)
			}
			user_data.UserGuest = GetCacheStatus("shutdown", roomId, uid)
			user_data.UserGuest = GetCacheStatus("songdown", roomId, uid)
			user_data.UserGuest = GetCacheStatus("guestctrl", roomId, uid)
			users = append(users, user_data)
		}
	}
	msgJson := make(map[string]interface{})
	msgJson["room_online"] = online
	err := models.UpdateRoom(roomId, msgJson)
	if err != nil {
		log.Println("更新房间信息失败")
	}
	msgJson = make(map[string]interface{})
	msgJson["type"] = "online"
	msgJson["channel"] = roomId
	msgJson["data"] = users
	SendMessageToRoom(roomId, msgJson)
}

func GetOnlineUserData(uid int) (OnlineUser, error) {
	data := OnlineUser{}
	u, err := models.GetUserInfoById(uid)
	if err != nil {
		return data, err
	}
	data = OnlineUser{
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
		UserAdmin:    true,
	}
	return data, nil
}

func GetCacheStatus(t string, roomId, uid int) bool {
	switch t {
	case "shutdown":
		if v, err := config.GetCacheString(fmt.Sprintf("shutdown_room_%d_user_%d", roomId, uid)); err == nil {
			if v == "true" {
				return true
			}
		}
	case "songdown":
		if v, err := config.GetCacheString(fmt.Sprintf("songdown_room_%d_user_%d", roomId, uid)); err == nil {
			if v == "true" {
				return true
			}
		}
		if v, err := config.GetCacheString(fmt.Sprintf("songdown_room_%d_user_%d", roomId, uid)); err == nil {
			if v == "true" {
				return true
			}
		}
	}
	return false
}

func SendMessageToRoom(roomId int, msgJson map[string]interface{}) {
	cs := GetClientsByRoom(roomId)
	for _, c := range cs {
		err := c.Socket.WriteJSON(msgJson)
		if err != nil {
			log.Printf("发送消息到房间:%d失败, %s", roomId, err)
		}
	}
}

func SendMsgToAll(msg map[string]interface{}) {
	for c, _ := range SCM.Clients {
		err := c.Socket.WriteJSON(msg)
		if err != nil {
			log.Printf("发送到用户：%s, 房间:%d失败:%s", c.UserId, c.RoomId, err)
		}
	}
}
