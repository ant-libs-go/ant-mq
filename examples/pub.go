/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-06 13:05:07
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"github.com/ant-libs-go/ant-mq/client"
	rpc "github.com/ant-libs-go/ant-mq/client/rpc"
)

func main() {
	options := []client.Option{
		client.WithServer("127.0.0.1:9472"),
		client.WithRetry(3),
	}
	c := rpc.New("ccclient", options...)
	c.Publish("aaa", []byte("textcontent"))
}
