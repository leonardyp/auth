package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/login","description":""}],"info":{"title":"hprose 之 tls","description":"使用hprose的tls rpc","contact":"1679550318@qq.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/login":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/login","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"post","type":"","summary":"Login","parameters":[{"paramType":"body","name":"login","description":"'用户登录信息json'","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"login","description":"'普通用户登陆信息'","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"'登录类型'","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":0,"message":"token","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
