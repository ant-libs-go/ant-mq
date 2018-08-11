/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-01 22:46:56
# File Name: handler.go
# Description:
####################################################################### */

package server

import (
	"log"

	"github.com/ant-libs-go/ant-mq/core"
	"github.com/ant-libs-go/ant-mq/core/types"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (this *Handler) Pub(req types.Request, resp *types.Response) (err error) {
	log.Printf("RECV|Topic-%s|Requester-%s", req.Header.Topic, req.Header.Requester)
	if err := core.Publish(req, resp); err != nil {
		log.Printf("RECV|FAIL|UUID-%s|PubExp|%s", resp.Header.SessionId, err)
		return err
	}
	log.Printf("RECV|SUCC|UUID-%s", resp.Header.SessionId)
	return nil
}
