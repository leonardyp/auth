package utils

import (
	"github.com/leonardyp/gocache"
	"time"
)

type TokenCache interface {
	//初始化数据结构
	CacheInit()
	// 获取缓存中的DataTable
	GetCacheTable() map[string]TokenUserInfo
	//判断令牌是否存在
	TokenIsExist(token string) bool
	//更新令牌过期时间
	TokenTimeUpdate(token string, dur time.Duration)
	// 添加令牌
	TokenInsert(token string, info interface{}, dur time.Duration)
	//token是否有效,用于校验用户是否更改过信息
	IsValid(token string)
}

type TokenCacheStruct struct {
	Cache map[string]TokenUserInfo
}

type TokenUserInfo struct {
	Expired  time.Duration `json:"expired"`
	UserInfo interface{}   `json:"userInfo"`
}

var tokenCacheStruct = &TokenCacheStruct{
	Cache: make(map[string]TokenUserInfo),
}

func GetAuth() *TokenCacheStruct {
	return tokenCacheStruct
}
func (this *TokenCacheStruct) CacheInit() {
	if this.Cache == nil {
		this.Cache = map[string]TokenUserInfo{}
	}
}
func (this *TokenCacheStruct) GetCacheTable() map[string]TokenUserInfo {
	this.CacheInit()
	return this.Cache
}
func (this *TokenCacheStruct) TokenIsExist(token string) bool {
	this.CacheInit()
	if _, ok := this.Cache[token]; !ok {
		return false
	} else {
		return true
	}
}

var cache = gocache.NewCache()

func (this *TokenCacheStruct) TokenTimeUpdate(token string, dur time.Duration) {
	this.CacheInit()
	if v, ok := this.Cache[token]; ok {
		v.Expired = dur
	}
}
func (this *TokenCacheStruct) TokenInsert(token string, info TokenUserInfo, dur time.Duration) {
	this.CacheInit()
	if !this.TokenIsExist(token) {
		this.Cache[token] = info
		//cache的过期时间为令牌过期时间*2
		cache.Set(token, this.Cache, this.Cache[token].Expired*2)
	} else {
		this.TokenTimeUpdate(token, dur)
	}
}
func (this *TokenCacheStruct) IsValid(token string) bool {
	return true
}
func (this *TokenCacheStruct) ClearToken(tokenValue string) {
	cache := this.GetCacheTable()
	if cache != nil {
		delete(cache, tokenValue)
	}
}
func (this *TokenCacheStruct) TokenGetCredence(tokenValue string) interface{} {
	if this.IsValid(tokenValue) {
		cache := this.GetCacheTable()
		if cache != nil {
			if v, ok := cache[tokenValue]; ok {
				return v.UserInfo
			}
		}
	}
	return nil
}
