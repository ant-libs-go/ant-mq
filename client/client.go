/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-02 18:32:31
# File Name: client.go
# Description:
####################################################################### */

package client

type Client interface {
	Publish(topic string, payload []byte) error
	Subscribe(topic string) (<-chan []byte, error)
	Close() error
}
