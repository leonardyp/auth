package main

import (
	_ "auth/docs"
	_ "auth/routers"
	_ "auth/rpc"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api"] = "swagger/api"
	}
	beego.Run()
}
