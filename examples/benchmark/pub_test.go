/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-06 15:05:07
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"fmt"
	"testing"

	"github.com/ant-libs-go/ant-mq/client"
	rpc "github.com/ant-libs-go/ant-mq/client/rpc"
)

func BenchmarkPub(b *testing.B) {

	b.ReportAllocs()

	options := []client.Option{
		client.WithServer("127.0.0.1:9472"),
		client.WithRetry(3),
	}
	c := rpc.New("ccclient", options...)

	for i := 0; i < b.N; i++ {
		c.Publish("aaa", []byte(fmt.Sprintf("test text %d", i)))
	}
}
