package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const heartbeatExpirationTime = 6 * 60

type Client struct {
	Addr          string
	Ticket        string
	Socket        *websocket.Conn
	UserId        string
	RoomId        int
	FirstInTime   uint64
	HeartbeatTime uint64
}

func NewClient(socket *websocket.Conn, addr, userId, ticket string, roomId int) {
	firstInTime := uint64(time.Now().Unix())
	client := &Client{
		Addr:          addr,
		Socket:        socket,
		UserId:        userId,
		RoomId:        roomId,
		FirstInTime:   firstInTime,
		HeartbeatTime: firstInTime,
		Ticket:        ticket,
	}
	SCM.Register <- client
	client.Socket.SetCloseHandler(func(code int, text string) error {
		SCM.Unregister <- client
		return nil
	})

	go func() {
		for {
			t, c, _ := client.Socket.ReadMessage()
			switch string(c) {
			case "bye":
				log.Printf("用户:%v(%s)房号:%d主动断开", client.UserId, client.Addr, client.RoomId)
				client.UnConnectSocket()
				SCM.Unregister <- client
				return
			}
			if t == -1 {
				return
			}
		}
	}()
	return
}
