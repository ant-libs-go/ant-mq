/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-06 13:05:07
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"fmt"
	"github.com/ant-libs-go/ant-mq/client"
	rpc "github.com/ant-libs-go/ant-mq/client/rpc"
	//"time"
)

func main() {
	options := []client.Option{
		client.WithListen("127.0.0.1:8085"),
	}
	c := rpc.New("ccclient", options...)

	handler := map[string]func(payload []byte) error{}

	handler["aaa"] = func(payload []byte) error {
		fmt.Println(fmt.Sprintf("get %s", payload))
		//time.Sleep(5 * time.Second)
		return nil
	}

	c.Subscribe(handler)
}