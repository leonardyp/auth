package models

type CommonResp struct {
	ErrorCode int         `json:"errorCode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}
