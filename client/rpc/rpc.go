/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-02 18:54:22
# File Name: rpc.go
# Description:
####################################################################### */

// TODO 还未考虑安全退出

package client

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/ant-libs-go/ant-mq/client"
)

type rpcClient struct {
	opts client.Options
}

var (
	rpcc        *rpc.Client
	listener    net.Listener
	subscribers = map[string]func(payload []byte) error{}
)

func publish(requester string, server string, topic string, payload []byte) (sessionId string, err error) {
	if rpcc == nil {
		rpcc, err = rpc.Dial("tcp", server)
		if err != nil {
			return "", fmt.Errorf("PUB|FAIL|ConnServerExp|%s", err)
		}
	}
	req := client.Request{}
	req.Header.Requester = requester
	req.Header.Topic = topic
	req.Body = payload
	resp := &client.Response{}
	err = rpcc.Call("AntMqServer.Pub", req, &resp)
	if err != nil {
		rpcc = nil
		return "", fmt.Errorf("PUB|FAIL|CallApiExp|%s", err)
	}
	if resp.Header.Status != client.RespStatusSucc {
		return "", fmt.Errorf("PUB|FAIL|CallApiExpStatus-%d", resp.Header.Status)
	}
	sessionId = resp.Header.SessionId
	return
}

func (this *rpcClient) Publish(topic string, payload []byte) (err error) {
	var sessionId string
	for i := int32(1); this.opts.Retry >= i; i++ {
		sessionId, err = publish(this.opts.Requester, this.opts.Server, topic, payload)
		if err == nil {
			log.Printf("PUB|SUCC-%s", sessionId)
			break
		}
		log.Printf("%s|Try-%d", err, i)
	}
	return
}

func (this *rpcClient) Subscribe(subs map[string]func(payload []byte) error) (err error) {
	for k, v := range subs {
		subscribers[k] = v
	}

	rpc.RegisterName("AntMqClient", NewHandler())
	listener, err = net.Listen("tcp", this.opts.Listen)
	if err != nil {
		return fmt.Errorf("SUB|FAIL|Listen-%s|%s", this.opts.Listen, err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
	return
}

func (this *rpcClient) Close() error {
	return nil
}

func New(name string, opts ...client.Option) *rpcClient {
	options := client.Options{
		Requester: name,
	}

	for _, opt := range opts {
		opt(&options)
	}

	o := &rpcClient{
		opts: options,
	}
	return o
}
