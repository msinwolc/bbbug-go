package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msinwolc/config"
	"github.com/msinwolc/middleware"
	"github.com/msinwolc/models"
	"github.com/msinwolc/util"
	"github.com/msinwolc/websocket"
)

type searchReq struct {
	BasicReq
	IsHots  int    `json:"isHots"`
	Keyword string `json:"keyword"`
	RoomId  int    `json:"room_id"`
}

type NowSong struct {
	User  UserInfo    `json:"user"`
	Song  models.Song `json:"song"`
	Since int64       `json:"since"`
}

func SearchSong(ctx *gin.Context) {
	var sr searchReq
	err := ctx.BindJSON(&sr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数不合法",
		})
		return
	}
	var songRs []models.SongResp
	if sr.IsHots == 1 {
		list, err := config.GetCacheString("week_song_play_rank")
		if err == nil {
			err := json.Unmarshal([]byte(list), &songRs)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"data": songRs,
				})
				return
			}
		}
		songRs, err := models.GetHotSongs()
		if err != nil {
			log.Println("获取热门歌曲失败")
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "获取热门歌曲失败",
			})
			return
		}
		songRsByte, err := json.Marshal(songRs)
		if err != nil {
			log.Printf("缓存热门歌曲失败, 错误:%s", err)
		}
		err = config.SetCacheString("week_song_play_rank", string(songRsByte), 30*time.Second)
		if err != nil {
			log.Printf("缓存热门歌曲失败, 错误:%s", err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": songRs,
		})
		return
	}
	rand.Seed(time.Now().UnixMilli())
	beginKeyWords := []string{"周杰伦", "林志炫", "梁静茹", "周华健", "伍佰", "五月天", "毛不易", "艾薇儿", "陈奕迅", "胡夏"}
	if sr.Keyword == "" {
		sr.Keyword = beginKeyWords[rand.Intn(len(beginKeyWords)-1)]
	}
	kwToken := rand.Int31n(99999999)
	client := &http.Client{}
	url := "http://bd.kuwo.cn/api/www/search/searchMusicBykeyWord?pn=1&rn=50&key=" + url.QueryEscape(sr.Keyword)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("查询歌曲失败:%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "搜索失败，请稍后重试",
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("csrf", fmt.Sprintf("%d", kwToken))
	req.Header.Set("Referer", "http://bd.kuwo.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win6e; x64) AppleWebKit/537.36 (KHTML, likeCecko) Chrome/99.0.4844.84 Safati/537.36")
	req.AddCookie(&http.Cookie{Name: "kw_token", Value: fmt.Sprintf("%d", kwToken)})
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("查询歌曲失败:%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "搜索失败，请稍后重试",
		})
		return
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("解析酷我:%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "搜索失败，请稍后重试",
		})
		return
	}
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respByte, &respMap)
	if err != nil {
		log.Printf("解析酷我:%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "搜索失败，请稍后重试",
		})
		return
	}
	if code, ok := respMap["code"]; ok {
		if code == 200.0 {
			if data, ok := respMap["data"]; ok {
				if list, ok := data.(map[string]interface{})["list"]; ok {
					for _, item := range list.([]interface{}) {
						var sr models.SongResp
						sr.Name = item.(map[string]interface{})["name"].(string)
						sr.Pic = item.(map[string]interface{})["pic"].(string)
						sr.Singer = item.(map[string]interface{})["artist"].(string)
						sr.Length = int(item.(map[string]interface{})["duration"].(float64))
						sr.Mid = int64(item.(map[string]interface{})["rid"].(float64))
						sr.Album = item.(map[string]interface{})["album"].(string)
						songRs = append(songRs, sr)
						srByte, err := json.Marshal(&sr)
						if err != nil {
							log.Printf("编码音乐信息失败:%s", err)
						}
						err = config.SetCacheString(fmt.Sprintf("song_detail_%d", sr.Mid), string(srByte), 3600*time.Second)
						if err != nil {
							log.Printf("缓存音乐信息失败:%s", err)
						}
					}
				}
			}
		}
	} else {
		log.Printf("获取歌曲信息失败:%s", respMap["message"])
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "搜索失败，请稍后重试",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": songRs,
	})
}

type operateSongReq struct {
	BasicReq
	RoomId int   `json:"room_id" binding:"required"`
	Mid    int64 `json:"mid"`
	At     int   `json:"at"`
}

type ListSong struct {
	User      UserInfo        `json:"user"`
	Song      models.SongResp `json:"song"`
	At        interface{}     `json:"at"`
	PushCount int             `json:"push_count"`
	PushTime  int64           `json:"push_time"`
	Since     int64           `json:"since"`
}

func AddSong(ctx *gin.Context) {
	var osr operateSongReq
	err := ctx.BindJSON(&osr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if osr.AccessToken == GuestToken {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(osr.AccessToken)
	if err != nil {
		log.Printf("用户信息获取失败:%s", err)
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
	r, err := models.GetRoomById(osr.RoomId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "房间信息获取失败",
		})
		return
	}
	if r.RoomType != 4 && r.RoomType != 1 {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "该房间不允许点歌",
		})
		return
	}
	if _, err := config.GetCacheString(fmt.Sprintf("songdown_room_%d_user_%d", osr.RoomId, uid)); err == nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "暂无点歌权限",
		})
		return
	}
	var sl ListSong
	sl.Song, err = GetSongDetailByMid(osr.Mid)
	if err != nil {
		log.Printf("歌曲信息获取失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "歌曲信息获取失败",
		})
		return
	}
	if sl.Song.Mid == 0 {
		s, err := models.GetSongByMid(osr.Mid)
		if err != nil {
			log.Printf("歌曲信息获取失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲数据获取失败",
			})
			return
		}
		sl.Song.Mid = s.SongMid
		sl.Song.Name = s.SongName
		sl.Song.Singer = s.SongSinger
		sl.Song.Pic = s.SongPic
		sl.Song.Length = s.SongLength
		sl.Song.Album = s.SongName
	}
	if sl.Song.Mid > 0 {
		gurl := fmt.Sprintf("http://wapi.kuwo.cn/api/www/music/musicInfo?mid=%d", osr.Mid)
		res, err := http.Get(gurl)
		if err != nil {
			log.Printf("获取歌曲信息失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲数据获取失败",
			})
			return
		}
		defer res.Body.Close()
		bodyByte, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("获取歌曲信息失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲数据获取失败",
			})
			return
		}
		body := make(map[string]interface{})
		err = json.Unmarshal(bodyByte, &body)
		if err != nil {
			log.Printf("获取歌曲信息失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲数据获取失败",
			})
			return
		}
		if code, ok := body["code"]; ok {
			if code == 200.0 {
				if data, ok := body["data"]; ok {
					if pic, ok := data.(map[string]interface{})["pic"]; ok {
						updateMap := make(map[string]interface{})
						updateMap["song_pic"] = pic.(string)
						err = models.UpdateSongByMid(osr.Mid, updateMap)
						if err != nil {
							log.Printf("更新歌曲信息失败:%s", err)
						}
						sl.Song.Pic = pic.(string)
						err = config.SetCacheString(fmt.Sprintf("song_picture_%d", osr.Mid), pic.(string), 0)
						if err != nil {
							log.Printf("更新歌曲信息失败:%s", err)
						}
						songByte, err := json.Marshal(sl.Song)
						if err == nil {
							err = config.SetCacheString(fmt.Sprintf("song_detail_%d", osr.Mid), string(songByte), 3600*time.Second)
							if err != nil {
								log.Printf("缓存歌曲信息失败:%s", err)
							}
						} else {
							log.Printf("更新歌曲信息失败:%s", err)
						}
					} else {
						log.Printf("歌曲信息获取失败, data为空:%v", data)
						ctx.JSON(200, gin.H{
							"code": 500,
							"msg":  "歌曲信息获取失败",
						})
						return
					}
				} else {
					log.Printf("歌曲信息获取失败:%s", code)
					ctx.JSON(200, gin.H{
						"code": 500,
						"msg":  "歌曲信息获取失败",
					})
					return
				}
			} else {
				log.Printf("歌曲信息获取失败:%s", code)
				ctx.JSON(200, gin.H{
					"code": 500,
					"msg":  "歌曲信息获取失败",
				})
				return
			}
		} else {
			log.Printf("歌曲信息获取失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲信息获取失败",
			})
			return
		}
	}
	if osr.At == 0 {
		sl.At = false
	} else {
		if sl.At == uid {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "错误",
			})
			return
		}
		atUser, err := GetUserData(osr.At)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "错误",
			})
			return
		}
		sl.At = atUser
	}
	sls := GetSongListFromCache(osr.RoomId)
	mysong := 0
	for _, sl := range sls {
		if sl.User.UserId == uid {
			mysong += 1
		}
		if sl.Song.Mid == osr.Mid {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("《%s》正在队列中", sl.Song.Name),
			})
			return
		}
	}
	addSongCdTime := r.RoomAddsongcd
	if !util.GetIsAdmin(u.UserGroup) && uid != r.RoomUser {
		if r.RoomAddsong == 1 {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "点歌失败,当前房间仅房主可点歌",
			})
			return
		}
		if addSongLastTimeStr, err := config.GetCacheString(fmt.Sprintf("song_add_%d_user_%d", osr.RoomId, uid)); err == nil {
			if addSongLastTime, err := strconv.ParseInt(addSongLastTimeStr, 10, 64); err == nil {
				addSongNeedTime := int64(addSongCdTime) - (time.Now().Unix() - addSongLastTime)
				if addSongNeedTime > 0 {
					ctx.JSON(200, gin.H{
						"code": 500,
						"msg":  fmt.Sprintf("点歌太频繁，请%ds后再试", addSongNeedTime),
					})
					return
				}
			}
		}
		if mysong > r.RoomAddcount {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("前面还有%d首歌没有播放, 急个锤子", mysong),
			})
			return
		}
	}
	if len(sls) == 1 && sls[0].User.UserId == 1 {
		sls = []ListSong{}
	}
	sl.User = u
	sls = append(sls, sl)
	slsByte, err := json.Marshal(sls)
	if err != nil {
		log.Printf("编码音乐列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "添加歌曲失败",
		})
		return
	}
	err = config.SetCacheString(fmt.Sprintf("SongList_%d", r.RoomId), string(slsByte), 86400*time.Second)
	if err != nil {
		log.Printf("缓存音乐列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "添加歌曲失败",
		})
		return
	}
	sendMsg := make(map[string]interface{})
	sendMsg["user"] = u
	sendMsg["song"] = sl.Song
	sendMsg["at"] = sl.At
	sendMsg["type"] = "addSong"
	sendMsg["count"] = len(sls)
	sendMsg["time"] = time.Now().Format("15:04:05")
	websocket.SendMessageToRoom(r.RoomId, sendMsg)
	os, err := models.GetSongByMidAndUser(osr.Mid, uid)
	if err != nil {
		ns := models.Song{
			SongSinger: sl.Song.Singer,
			SongPic:    sl.Song.Pic,
			SongMid:    sl.Song.Mid,
			SongName:   sl.Song.Name,
			SongLength: sl.Song.Length,
			SongUser:   uid,
			SongPlay:   1,
			SongWeek:   1,
		}
		_, err = models.CreateSong(ns)
		if err != nil {
			log.Printf("创建音乐信息失败:%s", err)
		}
	} else {
		updateMap := make(map[string]interface{})
		updateMap["song_play"] = os.SongPlay + 1
		updateMap["song_week"] = os.SongWeek + 1
		err = models.UpdateSongBySid(os.SongId, updateMap)
		if err != nil {
			log.Printf("更新音乐信息时失败:%s", err)
		}
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("《%s》已添加至歌单", sl.Song.Name),
		"data": sl.Song,
	})
}

func GetSongListFromCache(rid int) []ListSong {
	var ls []ListSong
	listStr, err := config.GetCacheString(fmt.Sprintf("SongList_%d", rid))
	if err == nil {
		err = json.Unmarshal([]byte(listStr), &ls)
		if err != nil {
			log.Printf("获取缓存歌曲信息失败:%s", err)
		}
	}
	return ls
}

func GetSongDetailByMid(mid int64) (models.SongResp, error) {
	var songR models.SongResp
	song, err := config.GetCacheString(fmt.Sprintf("song_detail_%d", mid))
	if err == nil {
		if err = json.Unmarshal([]byte(song), &songR); err == nil {
			if songR.Mid != 0 {
				return songR, nil
			}
		} else {
			log.Printf("获取缓存中的歌曲信息失败:%s", err)
		}
	}
	s, err := models.GetSongByMid(mid)
	if err != nil {
		return songR, fmt.Errorf("获取歌曲信息失败")
	}
	songR.Mid = s.SongMid
	songR.Name = s.SongName
	songR.Length = s.SongLength
	songR.Pic = s.SongPic
	songR.Singer = s.SongSinger
	return songR, nil
}

func PushSong(ctx *gin.Context) {
	var osr operateSongReq
	err := ctx.BindJSON(&osr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if osr.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(osr.AccessToken)
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
		log.Printf("获取用户信息失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户信息获取失败",
		})
		return
	}
	r, err := models.GetRoomById(osr.RoomId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取房间信息失败",
		})
		return
	}

	sls := GetSongListFromCache(osr.RoomId)
	push := false
	var pushSong ListSong
	pushIndex := 0
	pushRole := r.RoomUser != uid
	for i, l := range sls {
		if l.Song.Mid == osr.Mid {
			pushIndex = i
			pushSong = l
			push = true
			sls[i].PushTime = time.Now().Unix()
			if pushRole {
				sls[i].PushTime++
			} else {
				sls[i].PushTime = 888
			}
			break
		}
	}
	if !push {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "歌曲ID不存在",
		})
		return
	}
	limitCount := 0
	pushCacheTimeName := fmt.Sprintf("push_last_%d_%d", osr.RoomId, uid)
	if pushRole {
		pushCacheName := fmt.Sprintf("push_%d_%s_%d", osr.RoomId, time.Now().Format("2006-01-02"), uid)
		pushCount := r.RoomPushdaycount
		pushCache := 0
		var pushLastTime int64
		pushTimeLimit := r.RoomPushsongcd
		if pushCacheStr, err := config.GetCacheString(pushCacheName); err == nil {
			pushCache, _ = strconv.Atoi(pushCacheStr)
		}
		if pushLastTimeStr, err := config.GetCacheString(pushCacheTimeName); err == nil {
			pushLastTime, _ = strconv.ParseInt(pushLastTimeStr, 10, 60)
		}
		if pushCache > pushCount {
			if pushCount > 0 {
				ctx.JSON(200, gin.H{
					"code": 500,
					"msg":  fmt.Sprintf("你的%d次顶歌机会已使用完啦", pushCount),
				})
				return
			} else {
				ctx.JSON(200, gin.H{
					"code": 500,
					"msg":  "当前房间房主设置不允许顶歌",
				})
				return
			}
		}
		limitCount = pushCount - pushCache
		if time.Now().Unix()-pushLastTime < int64(pushTimeLimit) {
			timeStr := ""
			minute := math.Floor(float64(int64(pushTimeLimit)-(time.Now().Unix()-pushLastTime)) / 60)
			if minute > 0 {
				timeStr += fmt.Sprintf("%.0f分", minute)
			}
			second := (int64(pushTimeLimit) - (time.Now().Unix() - pushLastTime)) % 60
			timeStr += fmt.Sprintf("%d秒", second)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("顶歌太频繁啦，请%s后再试！", timeStr),
			})
			return
		}
		pushCache++
		pushCacheStr := strconv.Itoa(pushCache)
		err = config.SetCacheString(pushCacheName, pushCacheStr, 86400*time.Second)
		if err != nil {
			log.Printf("设置顶歌次数失败:%s", err)
		}
	}
	var nsls = []ListSong{pushSong}
	nsls = append(nsls, sls[:pushIndex]...)
	nsls = append(nsls, sls[pushIndex+1:]...)
	slsByte, err := json.Marshal(nsls)
	if err != nil {
		log.Printf("设置顶歌信息时失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "顶歌失败",
		})
		return
	}
	err = config.SetCacheString(fmt.Sprintf("SongList_%d", r.RoomId), string(slsByte), 86400*time.Second)
	if err != nil {
		log.Printf("缓存顶歌信息失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "顶歌失败",
		})
		return
	}
	sendMsg := make(map[string]interface{})
	sendMsg["user"] = u
	sendMsg["song"] = pushSong.Song
	sendMsg["type"] = "push"
	sendMsg["count"] = len(sls)
	sendMsg["time"] = time.Now().Format("15:04:05")
	websocket.SendMessageToRoom(r.RoomId, sendMsg)
	os, err := models.GetSongByMidAndUser(osr.Mid, uid)
	if err != nil {
		ns := models.Song{
			SongSinger: pushSong.Song.Singer,
			SongPic:    pushSong.Song.Pic,
			SongMid:    pushSong.Song.Mid,
			SongName:   pushSong.Song.Name,
			SongLength: pushSong.Song.Length,
			SongUser:   uid,
			SongPlay:   1,
			SongWeek:   1,
		}
		_, err := models.CreateSong(ns)
		if err != nil {
			log.Printf("创建音乐信息失败:%s", err)
		}
	} else {
		updateMap := make(map[string]interface{})
		updateMap["song_play"] = os.SongPlay + 1
		updateMap["song_week"] = os.SongWeek + 1
		err = models.UpdateSongBySid(os.SongId, updateMap)
		if err != nil {
			log.Printf("修改歌曲信息时失败:%s", err)
		}
	}
	err = config.SetCacheString(pushCacheTimeName, strconv.FormatInt(time.Now().Unix(), 10), 86400*time.Second)
	if err != nil {
		log.Printf("缓存点歌信息时失败:%s", err)
	}
	if pushRole {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf("顶歌成功，今日剩余%d", limitCount),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "点歌成功",
	})
}

func PlaySong(ctx *gin.Context) {
	var osr operateSongReq
	err := ctx.BindJSON(&osr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if osr.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(osr.AccessToken)
	if err != nil {
		log.Printf("获取用户信息失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "用户信息获取失败",
		})
		return
	}
	u, err := GetUserData(uid)
	if err != nil {
		log.Printf("获取用户信息失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "用户信息获取失败",
		})
		return
	}
	r, err := models.GetRoomById(osr.RoomId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "房间信息获取失败",
		})
	}
	if r.RoomType != 4 && r.RoomType != 1 {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "该房间不允许播放",
		})
		return
	}
	if !util.GetIsAdmin(uid) && r.RoomUser != uid {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "你没有播放权限！",
		})
		return
	}
	sls := GetSongListFromCache(osr.RoomId)
	isPush := false
	pushIndex := 0
	for i, sl := range sls {
		if sl.Song.Mid == osr.Mid {
			isPush = true
			pushIndex = i
			break
		}
	}
	if isPush {
		sls = append(sls[:pushIndex], sls[pushIndex+1:]...)
	}

	var sl ListSong
	sl.Song, err = GetSongDetailByMid(osr.Mid)
	if err != nil {
		log.Printf("歌曲信息获取失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "歌曲信息获取失败",
		})
		return
	}
	if sl.Song.Mid > 0 {
		gurl := fmt.Sprintf("http://wapi.kuwo.cn/api/www/music/musicInfo?mid=%d&httpsStatus=1", osr.Mid)
		res, err := http.Get(gurl)
		if err != nil {
			log.Printf("歌曲信息获取失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲信息获取失败",
			})
			return
		}
		defer res.Body.Close()
		bodyByte, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("读取歌曲信息失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲信息获取失败",
			})
			return
		}
		body := make(map[string]interface{})
		err = json.Unmarshal(bodyByte, &body)
		if err != nil {
			log.Printf("解析歌曲信息失败:%s", err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "歌曲信息获取失败",
			})
			return
		}
		if code, ok := body["code"]; ok {
			if code == 200.0 {
				if data, ok := body["data"]; ok {
					if pic, ok := data.(map[string]interface{})["pic"]; ok {
						updateMap := make(map[string]interface{})
						updateMap["song_pic"] = pic.(string)
						err = models.UpdateSongByMid(osr.Mid, updateMap)
						if err != nil {
							log.Printf("更新音乐信息失败:%s", err)
						}
						sl.Song.Pic = pic.(string)
						err = config.SetCacheString(fmt.Sprintf("song_picture_%d", osr.Mid), pic.(string), 0)
						if err != nil {
							log.Printf("更新音乐头像缓存失败:%s", err)
						}
						songByte, err := json.Marshal(sl.Song)
						if err == nil {
							err = config.SetCacheString(fmt.Sprintf("song_detail_%d", osr.Mid), string(songByte), 3600*time.Second)
							if err != nil {
								log.Printf("缓存歌曲信息失败:%s", err)
							}
						} else {
							log.Printf("编码音乐详情失败:%s", err)
						}
						return
					}
				}
			} else {
				log.Printf("歌曲信息获取失败:%s", code)
				ctx.JSON(200, gin.H{
					"code": 500,
					"msg":  "歌曲信息获取失败",
				})
				return
			}
		}
	}
	var nsls []ListSong

	sl.At = false
	sl.User = u
	nsls = append(nsls, sl)
	nsls = append(nsls, sls...)

	slsByte, err := json.Marshal(nsls)
	if err != nil {
		log.Printf("编码音乐列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放音乐失败",
		})
		return
	}
	err = config.SetCacheString(fmt.Sprintf("SongList_%d", r.RoomId), string(slsByte), 86400*time.Second)
	if err != nil {
		log.Printf("缓存音乐列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放音乐失败",
		})
		return
	}
	config.DeleCacheString(fmt.Sprintf("SongNow_%d", r.RoomId))
	os, err := models.GetSongByMidAndUser(osr.Mid, uid)
	if err != nil {
		ns := models.Song{
			SongSinger: sl.Song.Singer,
			SongPic:    sl.Song.Pic,
			SongMid:    sl.Song.Mid,
			SongName:   sl.Song.Name,
			SongLength: sl.Song.Length,
			SongUser:   uid,
			SongPlay:   1,
			SongWeek:   1,
		}
		_, err = models.CreateSong(ns)
		if err != nil {
			log.Printf("创建歌曲信息失败:%s", err)
		}
	} else {
		updateMap := make(map[string]interface{})
		updateMap["song_play"] = os.SongPlay
		updateMap["song_week"] = os.SongWeek
		err = models.UpdateSongBySid(os.SongId, updateMap)
		if err != nil {
			log.Printf("修改点歌信息失败:%s", err)
		}
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "播放成功",
	})
}

type lrcReq struct {
	BasicReq
	Mid int `json:"mid"`
}

type lrcResp struct {
	LineLyric string `json:"linelyric"`
	Time      string `json:"time"`
}

func PlayUrl(ctx *gin.Context) {
	var lr lrcReq
	err := ctx.BindJSON(&lr)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	cacheName := fmt.Sprintf("song_play_temp_url_%d", lr.Mid)
	url, err := config.GetCacheString(cacheName)
	if err == nil {
		ctx.Redirect(http.StatusMovedPermanently, url)
		return
	}
	if lr.Mid < 0 {
		a := ""
		ctx.Redirect(http.StatusMovedPermanently, a)
		return
	}
	gurl := fmt.Sprintf("http://antiserver.kuwo.cn/anti.s?type=convert_url&rid=%d&format=mp3&response=url", lr.Mid)
	res, err := http.Get(gurl)
	log.Println(res.Body)
	if err != nil {
		log.Printf("获取播放地址失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放地址获取失败",
		})
		return
	}
	defer res.Body.Close()
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("获取播放地址失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放地址获取失败",
		})
		return
	}
	// body := make(map[string]interface{})
	// err = json.Unmarshal(bodyByte, &body)
	// if err != nil {
	// 	log.Printf("获取播放地址失败:%s", err)
	// 	ctx.JSON(200, gin.H{
	// 		"code": 500,
	// 		"msg":  "播放地址获取失败",
	// 	})
	// 	return
	// }
	// if code, ok := body["code"]; ok {
	// 	if code == 200.0 {
	// 		if data, ok := body["data"]; ok {
	// 			if u, ok := data.(map[string]interface{})["url"]; ok {
	// 				err = config.SetCacheString(cacheName, u.(string), 30*time.Second)
	// 				if err != nil {
	// 					log.Printf("缓存播放路径出错:%s", err)
	// 				}
	ctx.Redirect(http.StatusMovedPermanently, string(bodyByte))
	// 			}
	// 		}
	// 	}
	// }
	// ctx.JSON(200, gin.H{
	// 	"code": 500,
	// 	"msg":  "获取播放地址失败",
	// })
}

func PlayUrlGet(ctx *gin.Context) {
	midStr := ctx.Query("mid")
	mid, err := strconv.ParseInt(midStr, 10, 64)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	cacheName := fmt.Sprintf("song_play_temp_url_%d", mid)
	url, err := config.GetCacheString(cacheName)
	if err == nil {
		ctx.Redirect(http.StatusMovedPermanently, url)
		return
	}
	if mid < 0 {
		// TODO
	}
	gurl := fmt.Sprintf("http://antiserver.kuwo.cn/anti.s?type=convert_url&rid=%d&format=mp3|acc&response=url", mid)
	res, err := http.Get(gurl)
	if err != nil {
		log.Printf("获取播放地址失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放地址获取失败",
		})
		return
	}
	defer res.Body.Close()
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("获取播放地址失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "播放地址获取失败",
		})
		return
	}
	// body := make(map[string]interface{})
	// err = json.Unmarshal(bodyByte, &body)
	// if err != nil {
	// 	log.Printf("获取播放地址失败:%s", err)
	// 	ctx.JSON(200, gin.H{
	// 		"code": 500,
	// 		"msg":  "播放地址获取失败",
	// 	})
	// 	return
	// }
	// log.Println(body)
	// if code, ok := body["code"]; ok {
	// 	if code == 200.0 {
	// 		if data, ok := body["data"]; ok {
	// 			if u, ok := data.(map[string]interface{})["url"]; ok {
	// 				err = config.SetCacheString(cacheName, u.(string), 30*time.Second)
	// 				if err != nil {
	// 					log.Printf("缓存播放路径出错:%s", err)
	// 				}
	log.Println(string(bodyByte))
	ctx.Redirect(http.StatusMovedPermanently, string(bodyByte))
	// 			}
	// 		}
	// 	}
	// }
	// ctx.JSON(200, gin.H{
	// 	"code": 500,
	// 	"msg":  "获取播放地址失败",
	// })
}

func GetLrc(ctx *gin.Context) {
	var lr lrcReq
	err := ctx.BindJSON(&lr)
	if err != nil {
		log.Printf("参数非法:%s", err)
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	data := make(map[string]interface{})
	if lr.Mid <= 0 {
		data["lineLyric"] = "歌曲为用户上传,暂无歌词"
		data["time"] = 0
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "歌曲为用户上传,暂无歌词",
			"data": data,
		})
		return
	}
	kwToken := rand.Int31n(99999999)
	client := &http.Client{}
	url := fmt.Sprintf("http://m.kuwo.cn/newh5/singles/songinfoandlrc?musicId=%d", lr.Mid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("查询歌词失败:%s", err)
		data["lineLyric"] = "没有查到歌词~"
		data["time"] = 0
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "没有查到歌词~",
			"data": data,
		})
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("csrf", fmt.Sprintf("%d", kwToken))
	req.Header.Set("Referer", "http://bd.kuwo.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	req.AddCookie(&http.Cookie{Name: "kw_token", Value: fmt.Sprintf("%d", kwToken)})
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("查询歌词失败:%s", err)
		data["lineLyric"] = "没有查到歌词~"
		data["time"] = 0
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "没有查到歌词~",
			"data": data,
		})
		return
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("解析歌词失败:%s", err)
		data["lineLyric"] = "很尴尬呀,没有查到歌词~"
		data["time"] = 0
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "很尴尬呀,没有查到歌词~",
			"data": data,
		})
		return
	}
	respMap := make(map[string]interface{})
	err = json.Unmarshal(respByte, &respMap)
	if err != nil {
		log.Printf("解析歌词Map失败:%s", err)
		data["lineLyric"] = "很尴尬呀,没有查到歌词~"
		data["time"] = 0
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "很尴尬呀,没有查到歌词~",
			"data": data,
		})
		return
	}
	lrs := []lrcResp{{LineLyric: "歌词加载成功", Time: "0"}}
	if status, ok := respMap["status"]; ok {
		if status == 200.0 {
			if data, ok := respMap["data"]; ok {
				if list, ok := data.(map[string]interface{})["lrclist"]; ok {
					for _, l := range list.([]interface{}) {
						lr := lrcResp{
							LineLyric: l.(map[string]interface{})["lineLyric"].(string),
							Time:      l.(map[string]interface{})["time"].(string),
						}
						lrs = append(lrs, lr)
					}
					ctx.JSON(200, gin.H{
						"code": 200,
						"msg":  "获取成功",
						"data": lrs,
					})
					return
				}
			}
		}
	}
	log.Printf("获取歌词响应失败:%s", respMap["msg"])
	data["lineLyric"] = "很尴尬呀,没有查到歌词~"
	data["time"] = 0
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "很尴尬呀,没有查到歌词~",
		"data": data,
	})
}

func RemoveSong(ctx *gin.Context) {
	var osr operateSongReq
	err := ctx.BindJSON(&osr)
	if err != nil {
		log.Printf("参数不合法:%s", err)
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数不合法",
		})
		return
	}
	if osr.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(osr.AccessToken)
	if err != nil {
		log.Printf("用户信息获取失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户信息获取失败",
		})
		return
	}
	u, err := GetUserData(uid)
	if err != nil {
		log.Printf("用户信息获取失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户信息获取失败",
		})
		return
	}
	r, err := models.GetRoomById(osr.RoomId)
	if err != nil {
		log.Printf("房间信息获取失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "房间信息获取失败",
		})
		return
	}
	ls := GetSongListFromCache(osr.RoomId)
	remove := false
	removeIndex := 0
	var removeSong ListSong
	for i, l := range ls {
		if l.Song.Mid == osr.Mid {
			removeSong = l
			removeIndex = i
			remove = true
			break
		}
	}
	if !remove {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "移除失败，歌曲ID不存在",
		})
		return
	}
	if r.RoomUser != uid && !util.GetIsAdmin(u.UserGroup) {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "你没有权限",
		})
		return
	}
	nls := append(ls[0:removeIndex], ls[removeIndex+1:]...)
	slsByte, err := json.Marshal(nls)
	if err != nil {
		log.Printf("编码待播放列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "移除音乐失败",
		})
		return
	}
	err = config.SetCacheString(fmt.Sprintf("SongList_%d", r.RoomId), string(slsByte), 86400*time.Second)
	if err != nil {
		log.Printf("缓存待播放列表失败:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "移除音乐失败",
		})
		return
	}
	sendMsg := make(map[string]interface{})
	sendMsg["user"] = u
	sendMsg["song"] = removeSong.Song
	sendMsg["type"] = "removeSong"
	sendMsg["count"] = len(ls)
	sendMsg["time"] = time.Now().Format("15:04:05")
	websocket.SendMessageToRoom(r.RoomId, sendMsg)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "移除成功",
	})
}

func PassSong(ctx *gin.Context) {
	var osr operateSongReq
	err := ctx.BindJSON(&osr)
	if err != nil {
		log.Printf("参数不合法:%s", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数不合法",
		})
		return
	}
	if osr.AccessToken == GuestToken {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	uid, err := middleware.GetUserId(osr.AccessToken)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	u, err := GetUserData(uid)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		return
	}
	r, err := models.GetRoomById(osr.RoomId)
	if err != nil {
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "获取房间信息失败",
			})
			return
		}
	}
	songResp, err := GetSongDetailByMid(osr.Mid)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "歌曲信息获取失败",
		})
		return
	}
	now, err := config.GetCacheString(fmt.Sprintf("SongNow_%d", r.RoomId))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "当前没有正在播放的歌曲",
		})
		return
	}
	var nowSong ListSong
	err = json.Unmarshal([]byte(now), &nowSong)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取正在播放的歌曲失败",
		})
		return
	}
	nowTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	err = config.SetCacheString(fmt.Sprintf("SongNextTime_%d", osr.RoomId), nowTimeStr, 0)
	if err != nil {
		log.Printf("设置下次播放时间失败:%s", err)
	}
	err = config.SetCacheString(fmt.Sprintf("SongNow_%d", osr.RoomId), "", 0)
	if err != nil {
		log.Printf("设置下次播放时间失败:%s", err)
	}
	sendMsg := make(map[string]interface{})
	sendMsg["user"] = u
	sendMsg["song"] = songResp
	sendMsg["type"] = "pass"
	sendMsg["time"] = time.Now().Format("15:04:05")
	websocket.SendMessageToRoom(osr.RoomId, sendMsg)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "切歌成功!",
	})
}
