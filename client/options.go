/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-02 18:46:45
# File Name: options.go
# Description:
####################################################################### */

package client

type Options struct {
	Requester string
	Server    string
	Listen    string
	Retry     int32
}

type Option func(o *Options)

func WithServer(inp string) Option {
	return func(o *Options) {
		o.Server = inp
	}
}

func WithListen(inp string) Option {
	return func(o *Options) {
		o.Listen = inp
	}
}

func WithRetry(inp int32) Option {
	return func(o *Options) {
		o.Retry = inp
	}
}
