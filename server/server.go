/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-01 20:28:48
# File Name: server.go
# Description:
####################################################################### */

package server

type Server interface {
	Run() error
	Close() error
}
