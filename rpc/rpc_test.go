package rpc

import (
	. "auth/rpc/tls"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-go"
	"io/ioutil"
	"testing"
	"time"
)

func client() {
	caCrt, _ := ioutil.ReadFile("tls/ca.crt")
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(caCrt))

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
	client.SetKeepAlive(true)
	a := stub.UserLogin("world", "demo")
	fmt.Println(a)
}

func TestRpc(t *testing.T) {
	time.Sleep(3 * time.Second)
	client()
}
