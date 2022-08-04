package models

type Room struct {
	RoomId           int         `gorm:"primarykey" json:"room_id" binding:"required"`
	RoomUser         int         `json:"room_user" `
	RoomAddsongcd    int         `json:"room_addsongcd" `
	RoomAddcount     int         `json:"room_addcount" `
	RoomPushdaycount int         `json:"room_pushdaycount" `
	RoomPushsongcd   int         `json:"room_pushsongcd" `
	RoomOnline       int         `json:"room_online"`
	RoomRealonline   int         `json:"room_realonline"`
	RoomHide         int         `json:"room_hide" `
	RoomName         string      `json:"room_name" binding:"required"`
	RoomType         int         `json:"room_type" `
	RoomPublic       int         `json:"room_public" `
	RoomPassword     string      `json:"room_password"`
	RoomNotice       string      `json:"room_notice"`
	RoomAddsong      int         `json:"room_addsong"`
	RoomSendmsg      int         `json:"room_sendmsg"`
	RoomRobot        int         `json:"room_robot" `
	RoomOrder        int         `json:"room_order" `
	RoomReason       string      `json:"room_reason"`
	RoomPlayone      int         `json:"room_playone" `
	RoomVotepass     int         `json:"room_votepass"`
	RoomVotepercent  int         `json:"room_votepercent"`
	RoomBackground   string      `json:"room_background"`
	RoomApp          string      `json:"room_app"`
	RoomFullpage     int         `json:"room_fullpage"`
	RoomStatus       int         `json:"room_status"`
	RoomCreatetime   int64       `gorm:"autoCreateTime" json:"room_createtime"`
	RoomUpdatetime   int64       `gorm:"autoUpdateTime" json:"room_updatetime"`
	Admin            interface{} `json:"admin" gorm:"-"`
}

func tableName() string {
	return "sa_room"
}

func GetRoomById(id int) (Room, error) {
	room := Room{}
	result := GetDB().Table(tableName()).Where("room_id = ?", id).First(&room)
	return room, result.Error
}

func GetRoomByUser(uid int) (Room, error) {
	room := Room{}
	result := GetDB().Table(tableName()).Where("room_user = ?", uid).First(&room)
	return room, result.Error
}

func UpdateRoom(roomId int, data map[string]interface{}) error {
	result := GetDB().Table(tableName()).Where("room_id = ?", roomId).Updates(data)
	return result.Error
}

func GetRooms() []Room {
	var rs []Room
	GetDB().Table("sa_room").Omit("room_createtime", "sa_user.user_head", "sa_user.user_group").
		Joins("left join sa_user on sa_room.room_user = sa_user.user_id").
		Where("(sa_room.room_online > 0 or sa_room.room_order > 10000) and sa_room.room_hide = 0").
		Order("sa_room.room_order desc, sa_room.room_online desc, sa_room.room_id desc").
		Find(&rs)
	return rs
}
