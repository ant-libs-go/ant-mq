/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-03 19:36:26
# File Name: director.go
# Description:
####################################################################### */

package director

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/naoina/toml"
)

var (
	_sr   *Source
	_lock sync.RWMutex
	once  sync.Once
)

type Topic struct {
	PersistFirst bool   // 持久化优先还是Pub优先
	PersistDb    string // 持久化方案，默认不持久化，格式：协议://info。可选协议file
}

type Subscriber struct {
	Addrs    []string
	Selector string // Pub策略，可选all,random
	Timeout  int32  // Pub超时时间，单位second
	Retry    int32  // 重试次数，默认为0，一直重试
}

// 订阅时，将Subscriber的信息拷贝到Relation中
type Relation struct {
	Topic         string
	Subscriber    string
	DownPersistDb string // 订阅者下线时是否持久化，默认不持久化，格式：协议://info。可选协议file
	Addrs         []string
	Selector      string // Pub策略，可选all,random
	Timeout       int32  // Pub超时时间，单位second
	Retry         int32  // 重试次数，默认为0，一直重试
}

type Source struct {
	Topic      map[string]*Topic
	Subscriber map[string]*Subscriber
	Relation   map[string]*Relation
}

func Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("BROKER|LoadSource|FAIL|OPEN|%s", err)
	}
	defer f.Close()

	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("BROKER|LoadSource|FAIL|READ|%s", err)
	}

	var r Source
	err = toml.Unmarshal(buff, &r)
	if err != nil {
		return fmt.Errorf("BROKER|LoadSource|FAIL|Unmarshal|%s", err)
	}
	_lock.Lock()
	_sr = &r
	_lock.Unlock()

	return nil
}

func Get(topic string) (topicInfo *Topic, relations []*Relation, err error) {
	_lock.RLock()
	defer _lock.RUnlock()

	topicInfo, ok := _sr.Topic[topic]
	if !ok {
		return nil, nil, fmt.Errorf("UnknownTopic-%s", topic)
	}
	for _, relation := range _sr.Relation {
		if relation.Topic != topic {
			continue
		}
		relations = append(relations, relation)
	}
	return
}
