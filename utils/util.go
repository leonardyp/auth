package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const (
	ERR_NO_ROWS_CODE         = 10000
	ERR_SINA_REQ_CODE        = 10001
	ERR_TOO_MUCH_PARAM_CODE  = 10008
	ERR_TOO_LESS_PARAM_CODE  = 10009
	QQ_THIRD_PART_TYPE       = 1
	WX_THIRD_PART_TYPE       = 2
	WB_THIRD_PART_TYPE       = 3
	CACHE_GROUP_PIC_TIME_OUT = 60 //s
	CITY_CODE_URL            = "http://open.weibo.com/wiki/%E7%9C%81%E4%BB%BD%E5%9F%8E%E5%B8%82%E7%BC%96%E7%A0%81%E8%A1%A8"
)

func Md5(bys []byte) string {
	h := md5.New()
	h.Write(bys)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//@Description base64加密
func Base64Encode(src []byte) []byte {
	return []byte(base64.URLEncoding.EncodeToString(src))
}

//@Description base64解密
func Base64Decode(src []byte) ([]byte, error) {
	return base64.URLEncoding.DecodeString(string(src))
}

//@Description 生成随机数
func RandomNum(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i := 0; i < length; i++ {
		str += fmt.Sprint(r.Intn(10))
	}
	return str
}

//@Description 时间磋和0-9999的随机数 的md5
func SessionId() string {
	now := time.Now().UnixNano()
	r := rand.New(rand.NewSource(now))
	randNum := r.Intn(10000)
	return Md5([]byte(fmt.Sprintf("%v%v", now, randNum)))
}

var seed = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"

// @Description 生成指定长度的字母和数字组成的字符串
func RandomStr(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i := 0; i < length; i++ {
		str += string(seed[r.Intn(62)])
	}
	return str
}

// @Description 生成指定长度的字母和数字组成的字符串
func RandomStrWithPrefix(prefix string, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i := 0; i < length; i++ {
		str += string(seed[r.Intn(62)])
	}
	return prefix + str
}
