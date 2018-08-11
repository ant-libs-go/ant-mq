/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-03 19:43:12
# File Name: publisher.go
# Description:
####################################################################### */

package publisher

import (
	"github.com/ant-libs-go/ant-mq/core/types"
)

type Publisher interface {
	Publish(addr string, req *types.Request, resp *types.Response) (r *Call, err error)
	Close() error
}

type Call struct {
	Done  chan int
	Error error
}

func New(transport string) (r Publisher) {
	switch transport {
	case "rpc":
		r = &rpcPublisher{}
	default:
		r = nil
	}
	return
}
