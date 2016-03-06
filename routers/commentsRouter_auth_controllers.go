package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["auth/controllers:LoginController"] = append(beego.GlobalControllerRouter["auth/controllers:LoginController"],
		beego.ControllerComments{
			"LoginAction",
			`/`,
			[]string{"post"},
			nil})

}
