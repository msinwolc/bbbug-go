package models

type Song struct {
	SongId         int `gorm:"primarykey"`
	SongUser       int
	SongMid        int64
	SongName       string
	SongSinger     string
	SongPic        string
	SongLength     int
	SongPlay       int
	SongWeek       int
	SongFav        int
	SongStatus     int
	SongCreatetime int64 `gorm:"autoCreateTime"`
	SongUpdatetime int64 `gorm:"autoUpdateTime"`
}

type SongResp struct {
	Id     int    `json:"id"`
	Mid    int64  `json:"mid"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Singer string `json:"singer"`
	Week   int    `json:"week"`
	Length int    `json:"length"`
	Album  string `json:"album"`
	Played int    `json:"played"`
}

func GetHotSongs() ([]SongResp, error) {
	var songRs []SongResp
	result := GetDB().Table("sa_song").Select("count(song_week) as week, song_mid as mid, max(song_id) as id, max(song_pic) as pic, max(song_singer) as singer, max(song_name) as name").Where("song_week > 0").Group("song_mid").Order("week desc").Limit(50).Find(&songRs)
	return songRs, result.Error
}

func GetSongByMidAndUser(mid int64, uid int) (Song, error) {
	var song Song
	result := GetDB().Table("sa_song").Where("song_mid = ? AND song_user = ?", mid, uid).First(&song)
	return song, result.Error
}

func GetSongByMid(mid int64) (Song, error) {
	var song Song
	result := GetDB().Table("sa_song").Where("song_mid = ?", mid).First(&song)
	return song, result.Error
}

func CreateSong(song Song) (Song, error) {
	result := GetDB().Table("sa_song").Omit("song_play", "song_week", "song_fav", "song_status").Create(&song)
	return song, result.Error
}

func DeleteSong(mid int64, uid int) error {
	result := GetDB().Table("sa_song").Where("song_mid = ? AND song_user = ?", mid, uid).Delete(&Song{})
	return result.Error
}

func GetSongList(uid, page, perpage int) ([]Song, error) {
	var songList []Song
	result := GetDB().Table("sa_song").Select("song_mid as mid,song_length as length,song_name as name,song_singer as singer,song_play as played,song_pic as pic").Where("song_user = ?",
		uid).Order("song_updatetime desc,song_play desc,song_id desc").Limit(perpage).Offset((page - 1) * perpage).Find(&songList)
	return songList, result.Error
}

func GetUserSongList(uid int, isAll bool) ([]SongResp, error) {
	var songRs []SongResp
	var err error
	if isAll {
		result := GetDB().Table("sa_song").Select("song_mid as mid,song_length as length,song_name as name,song_singer as singer,song_play as played,song_pic as pic").Where("song_user = ?",
			uid).Order("song_updatetime desc,song_play desc,song_id desc").Find(&songRs)
		err = result.Error
	} else {
		result := GetDB().Table("sa_song").Select("song_mid as mid,song_length as length,song_name as name,song_singer as singer,song_play as played,song_pic as pic").Where("song_user = ?",
			uid).Order("song_updatetime desc,song_play desc,song_id desc").Limit(50).Find(&songRs)
		err = result.Error
	}
	return songRs, err
}

func UpdateSongByMid(mid int64, column map[string]interface{}) error {
	result := GetDB().Table("sa_song").Where("song_mid = ?", mid).Updates(column)
	return result.Error
}

func UpdateSongBySid(sid int, column map[string]interface{}) error {
	result := GetDB().Table("sa_song").Where("song_id = ?", sid).Updates(column)
	return result.Error
}
