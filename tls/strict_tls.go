package crypto

import (
	"auth/logger"
	"io/ioutil"
)

func InitServerCrt(serverCrt string) ([]byte, error) {
	return ioutil.ReadFile(serverCrt)
}
func InitServerPrimaryKey(serverPrimaryKey string) ([]byte, error) {
	return ioutil.ReadFile(serverPrimaryKey)
}
func InitCaCrt(caCrt string) ([]byte, error) {
	return ioutil.ReadFile(caCrt)
}

var (
	ServerCrt,
	ServerPrimaryKey,
	CaCrt []byte
	err error
)

func init() {
	ServerCrt, err = InitServerCrt("tls/server.crt")
	if err != nil {
		logger.ErrorStd("%v", err)
		panic(err)
	}

	ServerPrimaryKey, err = InitServerPrimaryKey("tls/server.key")
	if err != nil {
		logger.ErrorStd("%v", err)
		panic(err)
	}
	CaCrt, err = InitCaCrt("tls/ca.crt")
	if err != nil {
		logger.ErrorStd("%v", err)
		panic(err)
	}
}
