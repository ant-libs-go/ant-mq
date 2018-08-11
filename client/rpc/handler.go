/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-02 18:56:37
# File Name: handler.go
# Description:
####################################################################### */

package client

import (
	"fmt"

	"github.com/ant-libs-go/ant-mq/client"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (this *Handler) Sub(req client.Request, resp *client.Response) (err error) {
	f := func(topic string) func(payload []byte) error {
		if fun, ok := subscribers[req.Header.Topic]; ok {
			return fun
		}
		return func(payload []byte) error {
			resp.Header.Status = client.RespStatusUnSub
			return fmt.Errorf("SUB|FAIL|UnknownSubscriber-%s", topic)
		}
	}

	resp.Header.Status = client.RespStatusSucc

	err = f(req.Header.Topic)(req.Body)
	if err == nil {
		return
	}
	if resp.Header.Status == client.RespStatusUnSub {
		return fmt.Errorf("UnknownSubscriber-%s", req.Header.Topic)
	}
	resp.Header.Status = client.RespStatusFail
	resp.Body = []byte(fmt.Sprintf("SubscriberErr, %s", err))
	return nil
}
