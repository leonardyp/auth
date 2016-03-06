package rpc

import (
	"auth/logger"
	"auth/models"
	. "auth/rpc/tls"
	"auth/utils"
	"crypto/tls"
	"crypto/x509"
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-go"
)

type ClientStub struct {
	UserLogin func(reqStr, remoteAddr string) models.CommonResp
}

func GetClientStub() *ClientStub {
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(CaCrt))

	cliCrt, _ := tls.LoadX509KeyPair("tls/client.crt", "tls/client.key")

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(CaCrt))

	client := hprose.NewClient("tcp4://" + beego.AppConfig.String("rpc::RPCADDR") + ":" + beego.AppConfig.String("rpc::RPCPORT"))
	var stub *ClientStub
	client.UseService(&stub)
	client.SetTLSClientConfig(&tls.Config{
		RootCAs:            roots,
		InsecureSkipVerify: false,
		ServerName:         "server",
		Certificates:       []tls.Certificate{cliCrt},
	})
	return stub
}
func init() {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(CaCrt))

	config := &tls.Config{
		ClientCAs:  pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	config.NextProtos = []string{"http/2.0"}

	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], _ = tls.X509KeyPair(append(ServerCrt, ServerPrimaryKey...), append(ServerCrt, ServerPrimaryKey...))

	server := hprose.NewTcpServer("tcp://" + beego.AppConfig.String("rpc::RPCADDR") + ":" + beego.AppConfig.String("rpc::RPCPORT"))
	server.SetTLSConfig(config)

	server.AddFunction("UserLogin", UserLogin, true)

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()
}

func UserLogin(reqStr string, remoteAddr string) models.CommonResp {
	logger.DebugStd("rpc param:...%v:%v", reqStr, remoteAddr)
	if len(reqStr) < 0 {
		logger.Error("%v", "illegal request param length")
		return models.CommonResp{
			ErrorCode: -1,
			Msg:       "请求参数:" + reqStr + "非法",
		}
	} else {
		session := utils.SessionId()
		//todo 添加登录逻辑
		return models.CommonResp{
			Data: "登录成功数据:" + session,
		}
	}
}
