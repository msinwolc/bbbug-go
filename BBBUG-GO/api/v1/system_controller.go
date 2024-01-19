package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msinwolc/config"
	"github.com/msinwolc/models"
	"github.com/msinwolc/websocket"
)

const GuestToken = "45af3cfe44942c956e026d5fd58f0feffbd3a237"

type BasicReq struct {
	Plat        string `json:"plat" binding:"required"`
	Version     int    `json:"version" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}

func GetTime(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["time"] = time.Now().UnixNano() / 1e6
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func InitListener() {
	go Listener()
}

func Listener() {
	for {
		time.Sleep(500 * time.Millisecond)
		rs := getRoomList()
		if len(rs) == 0 {
			continue
		}
		for _, r := range rs {
			if s, err := getPlaySongByRid(r.RoomId); err == nil {
				if s.Song.Mid != 0 {
					if time.Now().Unix() < int64(s.Song.Length)+s.Since {
						preLoadMusicUrl(r)
						continue
					}
					if r.RoomType == 4 && r.RoomPlayone == 1 {
						playSong(r.RoomId, s, true)
						continue
					}
				}
			}
			if s, err := getSongFromList(r.RoomId); err == nil {
				if s.Song.Mid != 0 {
					playSong(r.RoomId, s, false)
					continue
				}
			}
			if r.RoomType == 4 {
				if s, err := getSongByUser(r.RoomUser); err == nil {
					playSong(r.RoomId, s, false)
				}
			} else {
				if r.RoomRobot == 0 {
					if s, err := getSongByRobot(); err == nil {
						playSong(r.RoomId, s, false)
					}
				}
			}
		}
	}
}

func getRoomList() []models.Room {
	if roomStr, err := config.GetCacheString("RoomList"); err == nil {
		var rs []models.Room
		err := json.Unmarshal([]byte(roomStr), &rs)
		if err == nil {
			return rs
		}
	}
	rs := models.GetRooms()
	roomByte, err := json.Marshal(rs)
	if err != nil {
		log.Printf("编码在线房间失败:%s", err)
	}
	err = config.SetCacheString("RoomList", string(roomByte), 5*time.Second)
	if err != nil {
		log.Printf("缓存在线房间失败:%s", err)
	}
	return rs
}

func getSongFromList(rid int) (ListSong, error) {
	var lss []ListSong
	var ls ListSong
	if lssStr, err := config.GetCacheString(fmt.Sprintf("SongList_%d", rid)); err == nil {
		err = json.Unmarshal([]byte(lssStr), &lss)
		if err != nil {
			log.Printf("解析待播放歌曲失败:%s", err)
			return ls, fmt.Errorf("解析待播放歌曲失败:%s", err)
		}
		if len(lss) > 0 {
			ls = lss[0]
			ls.Since = time.Now().Unix() + 5
			lss = lss[1:]
			lssByte, err := json.Marshal(lss)
			if err != nil {
				log.Printf("编码待播放歌曲失败:%s", err)
			} else {
				err = config.SetCacheString(fmt.Sprintf("SongList_%d", rid), string(lssByte), 86400*time.Second)
				if err != nil {
					log.Printf("编码待播放歌曲失败:%s", err)
				}
			}
		}
		return ls, nil
	}
	return ls, fmt.Errorf("未找到音乐")
}

func getPlaySongByRid(rid int) (ListSong, error) {
	var ls ListSong
	if nowStr, err := config.GetCacheString(fmt.Sprintf("SongNow_%d", rid)); err == nil {
		err = json.Unmarshal([]byte(nowStr), &ls)
		if err != nil {
			log.Printf("解析正在播放的音乐失败:%s", err)
			return ls, err
		}
		return ls, nil
	}
	return ls, fmt.Errorf("未找到音乐")
}

func preLoadMusicUrl(r models.Room) {
	sls := GetSongListFromCache(r.RoomId)
	have := false
	var now ListSong
	if len(sls) > 0 {
		return
	} else {
		if r.RoomType == 4 {
			if s, err := getSongByUser(r.RoomUser); err == nil {
				now = s
				have = true
			}
		} else {
			if r.RoomRobot == 0 {
				if s, err := getSongByRobot(); err == nil {
					now = s
					have = true
				}
			}
		}
		if have {
			addSongToList(r.RoomId, now)
		}
	}
}

func getSongByUser(uid int) (ListSong, error) {
	var ls ListSong
	u, err := GetUserData(uid)
	if err != nil {
		return ls, fmt.Errorf("未找到用户:%s", err)
	}
	ls.User = u
	sl, err := models.GetUserSongList(uid, true)
	if err != nil {
		return ls, fmt.Errorf("未找到该用户的音乐:%s", err)
	}
	if len(sl) > 1 {
		rand.Seed(time.Now().UnixMilli())
		s := sl[rand.Int31n(int32(len(sl))-1)]
		ls.Song = s
		ls.Since = time.Now().Unix()
	}
	return ls, nil
}

// TODO
func getSongByRobot() (ListSong, error) {
	return ListSong{}, nil
}

func addSongToList(rid int, ls ListSong) {
	sls := GetSongListFromCache(rid)
	isExist := false
	for _, sl := range sls {
		if sl.Song.Mid == ls.Song.Mid {
			isExist = true
			break
		}
	}
	if !isExist {
		sls = append(sls, ls)
		slsByte, err := json.Marshal(sls)
		if err != nil {
			log.Printf("编码待播放列表失败:%s", err)
		}
		err = config.SetCacheString(fmt.Sprintf("SongList_%d", rid), string(slsByte), 86400*time.Second)
		if err != nil {
			log.Printf("缓存待播放列表失败:%s", err)
		}
	}
}

func playSong(rid int, sl ListSong, save bool) {
	sl.Since = time.Now().Unix()
	nowCacheName := fmt.Sprintf("SongNow_%d", rid)
	slByte, err := json.Marshal(sl)
	if err != nil {
		log.Printf("编码正在播放的音乐失败:%s", err)
	} else {
		if save {
			err = config.SetCacheString(nowCacheName, string(slByte), 0)
			if err != nil {
				log.Printf("缓存正在播放的歌曲失败:%s", err)
			}
		} else {
			err = config.SetCacheString(nowCacheName, string(slByte), 3600*time.Second)
			if err != nil {
				log.Printf("缓存正在播放的歌曲失败:%s", err)
			}
		}
		songByte, err := json.Marshal(sl.Song)
		if err == nil {
			err = config.SetCacheString(fmt.Sprintf("song_detail_%d", sl.Song.Mid), string(songByte), 3600*time.Second)
			if err != nil {
				log.Printf("缓存正在播放的歌曲失败:%s", err)
			}
		} else {
			log.Printf("缓存正在播放的歌曲失败:%s", err)
		}
	}
	sendMsg := make(map[string]interface{})
	if sl.At == false {
		atMap := make(map[string]interface{})
		atMap["user_id"] = 0
		sl.At = atMap
	}
	sendMsg["at"] = sl.At
	sendMsg["user"] = sl.User
	sendMsg["song"] = sl.Song
	sendMsg["type"] = "playSong"
	sendMsg["since"] = time.Now().Unix()
	sendMsg["count"] = len(GetSongListFromCache(rid))
	sendMsg["time"] = time.Now().Format("15:04:05")
	websocket.SendMessageToRoom(rid, sendMsg)
}
