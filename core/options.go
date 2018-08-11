/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-04 16:54:00
# File Name: options.go
# Description:
####################################################################### */

package core

type Options struct {
	SourceFile        string
	PubGoroutineLimit int
	Transport         string
}

type Option func(o *Options)

func WithSourceFile(inp string) Option {
	return func(o *Options) {
		o.SourceFile = inp
	}
}

func WithPubGoroutineLimit(inp int) Option {
	return func(o *Options) {
		o.PubGoroutineLimit = inp
	}
}

func WithTransport(inp string) Option {
	return func(o *Options) {
		o.Transport = inp
	}
}
