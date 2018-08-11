/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-05 20:59:38
# File Name: random.go
# Description:
####################################################################### */

package selector

import (
	"time"
)

var (
	_randomInstance *randomSelector
)

type randomSelector struct{}

func NewRandomSelector() *randomSelector {
	once.Do(func() {
		_randomInstance = &randomSelector{}
	})
	return _randomInstance
}

func (this *randomSelector) Get(addrs []string) (r []string) {
	if len(r) > 1 {
		addrs = []string{addrs[time.Now().Unix()%int64(len(addrs))]}
	}
	return addrs
}
