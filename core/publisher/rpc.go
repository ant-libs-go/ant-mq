/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-03 19:42:00
# File Name: rpc.go
# Description:
####################################################################### */

// TODO 这里后续实现连接池

package publisher

import (
	"fmt"
	"net/rpc"

	"github.com/ant-libs-go/ant-mq/core/types"
)

var (
	rpcc *rpc.Client
)

type rpcPublisher struct{}

func (this *rpcPublisher) Publish(addr string, req *types.Request, resp *types.Response) (r *Call, err error) {
	r = &Call{Done: make(chan int)}

	rpcc, err = rpc.Dial("tcp", addr)
	if err != nil {
		err = fmt.Errorf("ConnServerExp-%s", err)
		return
	}

	go func() {
		err = rpcc.Call("AntMqClient.Sub", req, &resp)
		if err != nil {
			r.Error = err
		}
		r.Done <- 1
	}()
	return
}

func (this *rpcPublisher) Close() error {
	rpcc.Close()
	return nil
}
