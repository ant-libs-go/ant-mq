/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-01 20:27:06
# File Name: options.go
# Description:
####################################################################### */

package server

type Options struct {
	Listen string
}

type Option func(o *Options)

func WithListen(inp string) Option {
	return func(o *Options) {
		o.Listen = inp
	}
}
