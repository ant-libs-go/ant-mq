/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-01 22:36:57
# File Name: rpc.go
# Description:
####################################################################### */

package server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/ant-libs-go/ant-mq/server"
)

type rpcServer struct {
	opts server.Options
}

func (this *rpcServer) Run() (err error) {
	rpc.RegisterName("AntMqServer", NewHandler())
	listener, err := net.Listen("tcp", this.opts.Listen)
	if err != nil {
		return fmt.Errorf("RUN|FAIL|Listen-%s|%s", this.opts.Listen, err)
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

func (this *rpcServer) Close() error {
	return nil
}

func New(opts ...server.Option) *rpcServer {
	options := server.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	o := &rpcServer{
		opts: options,
	}
	return o
}
