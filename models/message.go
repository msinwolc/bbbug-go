package models

type Message struct {
	MessageId         int64  `gorm:"primarykey" json:"message_id"`
	MessageUser       int    `json:"message_user"`
	MessageType       string `json:"message_type"`
	MessageWhere      string `json:"message_where"`
	MessageTo         string `json:"message_to"`
	MessageContent    string `json:"message_content"`
	MessageStatus     int    `json:"message_status"`
	MessageCreatetime int64  `gorm:"autoCreateTime" json:"message_createtime"`
	MessageUpdatetime int64  `gorm:"autoUpdateTime" json:"message_updatetime"`
}

func GetMessageById(mid int) (Message, error) {
	msg := Message{}
	result := GetDB().Table("sa_message").Where("message_id = ?", mid).First(&msg)
	return msg, result.Error
}

func GetMessageListByMap(m map[string]interface{}, page, limit int) []Message {
	var list []Message
	GetDB().Table("sa_message").Where(m).Offset((page - 1) * limit).Limit(page).Find(&list)
	return list
}

func CreateMessage(m Message) (Message, error) {
	result := GetDB().Table("sa_message").Create(&m)
	return m, result.Error
}

func UpdateMessageById(id int64, column map[string]interface{}) error {
	result := GetDB().Table("sa_message").Where("message_id = ?", id).Updates(column)
	return result.Error
}

func DeleteMessageByRid(rid string) error {
	return GetDB().Table("sa_message").Where("message_to = ?", rid).Where("message_where = ?", "channel").Delete(&Message{}).Error
}

func DeleteMessageById(id int64) error {
	return GetDB().Table("sa_message").Where("message_id = ?", id).Delete(&Message{}).Error
}
