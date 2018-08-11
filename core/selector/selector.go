/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-05 20:38:56
# File Name: selector.go
# Description:
####################################################################### */

package selector

import (
	"sync"
)

var (
	once sync.Once
)

type Selector interface {
	Get(addrs []string) (r []string)
}

func New(sel string) (r Selector) {
	switch sel {
	case "random":
		r = NewRandomSelector()
	case "all":
		r = NewAllSelector()
	default:
		r = NewAllSelector()
	}
	return
}
