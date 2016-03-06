package controllers

import (
	"auth/logger"
	"auth/logic"
	"auth/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// @Title post
// @Description Login
// @Param 	login	 	body 	string 	 false   '用户登录信息json'
// @Param	login		query 	string	 false	 '普通用户登陆信息'
// @Param   type        query   string   true    '登录类型'
// @Success token
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) LoginAction() {
	requestBody := this.Ctx.Input.RequestBody
	typeStr := this.GetString("type")
	login := this.GetString("login")
	logger.DebugStd("req:....%v:%v:%v", string(requestBody), typeStr, login)
	if typeStr == "0" {
		if login == "" {
			this.Data["json"] = models.CommonResp{
				ErrorCode: -1,
				Msg:       "login  信息不能为空",
			}
			this.ServeJSON()
			return
		}
	} else if typeStr == "1" || typeStr == "2" || typeStr == "3" {
		if string(requestBody) == "" {
			this.Data["json"] = models.CommonResp{
				ErrorCode: -1,
				Msg:       "login体,信息不能为空",
			}
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = models.CommonResp{
			ErrorCode: -1,
			Msg:       "非法登录类型",
		}
		this.ServeJSON()
		return
	}

	_, userInfo, err := logic.LoginLogic(requestBody, typeStr, login, this.Ctx.Request.RemoteAddr)
	if err != nil {
		this.Data["json"] = models.CommonResp{
			ErrorCode: -1,
			Msg:       err.Error(),
		}
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = models.CommonResp{
			Data: userInfo,
		}
		this.ServeJSON()
		return
	}
}
