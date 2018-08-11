/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-05 20:56:01
# File Name: all.go
# Description:
####################################################################### */

package selector

import ()

var (
	_allInstance *allSelector
)

type allSelector struct{}

func NewAllSelector() *allSelector {
	once.Do(func() {
		_allInstance = &allSelector{}
	})
	return _allInstance
}

func (this *allSelector) Get(addrs []string) (r []string) {
	return addrs
}
