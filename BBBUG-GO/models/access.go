package models

type Access struct {
	AccessId         int `gorm:"primarykey"`
	AccessUser       int
	AccessToken      string
	AccessPlat       string
	AccessIp         string
	AccessStatus     int
	AccessCreatetime int64
	AccessUpdatetime int64
}

func CreateAccess(userId int, plat, token, ip string) error {
	acc := Access{
		AccessUser:  userId,
		AccessPlat:  plat,
		AccessToken: token,
		AccessIp:    ip,
	}
	result := GetDB().Table("sa_access").Create(&acc)
	return result.Error
}
