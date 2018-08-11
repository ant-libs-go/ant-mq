/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-02 18:28:02
# File Name: types.go
# Description:
####################################################################### */

package client

// 必须与client中的定义一致
type resp_status int

const (
	RespStatusSucc    resp_status = 200
	RespStatusUnTopic resp_status = 401
	RespStatusUnSub   resp_status = 404
	RespStatusFail    resp_status = 500
)

type Request struct {
	Header struct {
		Requester string
		Topic     string
	}
	Body []byte
}

type Response struct {
	Header struct {
		SessionId string
		Status    resp_status
	}
	Body []byte
}
