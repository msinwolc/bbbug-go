package models

import "time"

type Access struct {
	AccessId         int `gorm:"primarykey"`
	AccessUser       int
	AccessToken      string
	AccessPlat       string
	AccessIp         string
	AccessStatus     int
	AccessCreatetime time.Time
	AccessUpdatetime time.Time
}

func CreateAccess(userId int, plat, token, ip string) error {
	acc := Access{
		AccessUser:       userId,
		AccessPlat:       plat,
		AccessToken:      token,
		AccessIp:         ip,
		AccessCreatetime: time.Now(),
		AccessUpdatetime: time.Now(),
	}
	result := GetDB().Table("sa_access").Create(&acc)
	return result.Error
}
