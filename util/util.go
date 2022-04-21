package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeBySalt(password, salt string) string {
	m := md5.New()
	m.Write([]byte(password))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}
