package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func EncodeBySalt(password, salt string) string {
	m := md5.New()
	m.Write([]byte(password))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

func GetAddrPath(addr string) string {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", "http://whois.pcoline.com.cn/ip.jsp?ip="+addr, nil)
	if err != nil {
		return "未知"
	}
	res, err := cli.Do(req)
	if err != nil {
		return "未知"
	}
	defer res.Body.Close()
	data := transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(data)
	if err != nil {
		log.Println("transform to GBK error")
		return "未知"
	}
	out := strings.Replace(string(body), "\n", "", -1)
	out = strings.Replace(string(out), "\r", "", -1)
	return out
}

func GetWebsocketTicket(ip string, channel int) string {
	old := fmt.Sprintf("account%schannel%vsalt%v", ip, channel, channel)
	return EncodeBySalt(old, "socket")
}

func GetIsAdmin(group int) bool {
	var adminGroups = []int{1}
	for _, ag := range adminGroups {
		if ag == group {
			return true
		}
	}
	return false
}
