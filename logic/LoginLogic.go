package logic

import (
	"auth/logger"
	"auth/rpc"
	"auth/utils"
	"errors"
	"time"
)

var tokenCacheStruct = utils.GetAuth()

func LoginLogic(requestBody []byte, typeStr, login, remoteAddr string) (string, utils.TokenUserInfo, error) {
	var isLogin bool
	var token string
	var tokenUserInfo = utils.TokenUserInfo{}
	var ro = rpc.GetClientStub()
	switch typeStr {
	case "0": //普通,可以添加其他登录类型
		resp := ro.UserLogin(login, remoteAddr)
		if resp.ErrorCode != 0 {
			logger.DebugStd("2....code:%v,msg:%v", resp.ErrorCode, resp.Msg)
		} else {
			isLogin = true
			tokenUserInfo.UserInfo = resp.Data
		}
	default:
		return token, tokenUserInfo, errors.New("无法识别的登录类型:" + typeStr)
	}
	if isLogin {
		//产生令牌
		token = utils.GenGUID()
		if !tokenCacheStruct.TokenIsExist(token) {
			tokenCacheStruct.TokenInsert(token, tokenUserInfo, 1*time.Hour)
		}
		return token, tokenUserInfo, nil
	} else {
		return token, tokenUserInfo, errors.New("抱歉，登录失败")
	}
}
