/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-04 16:58:45
# File Name: broker.go
# Description:
####################################################################### */

package core

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ant-libs-go/ant-mq/core/director"
	"github.com/ant-libs-go/ant-mq/core/publisher"
	"github.com/ant-libs-go/ant-mq/core/selector"
	"github.com/ant-libs-go/ant-mq/core/types"
	"github.com/satori/go.uuid"
)

var (
	DefaultBroker *broker = NewBroker()
)

type broker struct {
	opts         Options
	GoroutineCnt chan int
}

func (this *broker) Publish(req types.Request, resp *types.Response) error {
	resp.Header.Status = types.RespStatusSucc

	sessionId, err := uuid.NewV4()
	if err != nil {
		resp.Header.Status = types.RespStatusFail
		resp.Header.SessionId = "Unknown"
		return fmt.Errorf("GenUUID|%s", err)
	}
	resp.Header.SessionId = sessionId.String()

	_, relations, err := director.Get(req.Header.Topic)
	if err != nil {
		resp.Header.Status = types.RespStatusUnTopic
		return err
	}

	// TODO 判断topic层是否要持久化

	// 后面的错误就不对pub client暴露了
	for _, relation := range relations {
		this.GoroutineCnt <- 1
		go func(relation *director.Relation) {
			log.Printf("PUB|UUID-%s|Subscriber-%s", resp.Header.SessionId, relation.Subscriber)
			err := this.publish(req.Header.Topic, req.Body, relation)
			if err == nil {
				log.Printf("PUB|SUCC|UUID-%s|Subscriber-%s", resp.Header.SessionId, relation.Subscriber)
				return
			}
			log.Printf("PUB|FAIL|UUID-%s|Subscriber-%s|PubExp|%s", resp.Header.SessionId, relation.Subscriber, err)

			// TODO 失败的时候是否需要持久化
		}(relation)
	}

	return nil
}

func (this *broker) publish(topic string, payload []byte, relation *director.Relation) (err error) {
	defer func() {
		<-this.GoroutineCnt
		if e := recover(); e != nil {
			err = fmt.Errorf("PublishUnknownExp-%s", e)
		}
	}()

	f := func(addr string) error {
		req := &types.Request{}
		req.Header.Requester = "ant-mq"
		req.Header.Topic = topic
		req.Body = payload
		resp := &types.Response{}

		puber := publisher.New(this.opts.Transport)
		call, err := puber.Publish(addr, req, resp)
		if err != nil {
			return err
		}

		// 统一采用这里的timeout机制
		select {
		case <-time.After(time.Duration(relation.Timeout) * time.Second):
			puber.Close()
			return fmt.Errorf("CallApiTimeout")
		case <-call.Done:
			if call.Error != nil {
				return fmt.Errorf("CallApiExp-%s", call.Error)
			}
			if resp.Header.Status != types.RespStatusSucc {
				return fmt.Errorf("CallApiExpStatus-%d", resp.Header.Status)
			}
		}
		return nil
	}

	// 只要一个地址失败则全标记为失败
	for _, addr := range selector.New(relation.Selector).Get(relation.Addrs) {
		var err error
		for i := int32(1); relation.Retry >= i || relation.Retry == 0; i++ {
			err = f(addr)
			if err == nil {
				break
			}
			time.Sleep(time.Second * 2)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBroker(opts ...Option) *broker {
	options := Options{
		SourceFile:        fmt.Sprintf("%s/%s", os.Getenv("WORKDIR"), "/runtime/source.toml"),
		PubGoroutineLimit: 10,
	}

	for _, opt := range opts {
		opt(&options)
	}

	o := &broker{
		opts:         options,
		GoroutineCnt: make(chan int, options.PubGoroutineLimit),
	}
	return o
}

func Publish(req types.Request, resp *types.Response) error {
	return DefaultBroker.Publish(req, resp)
}

func LoadSubscribeRelation() error {
	return director.Load(DefaultBroker.opts.SourceFile)
}
