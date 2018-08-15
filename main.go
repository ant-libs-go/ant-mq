/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-07-31 13:37:17
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"flag"
	"log"
	"os"

	"github.com/ant-libs-go/ant-mq/core"
	"github.com/ant-libs-go/ant-mq/server"
	rpcsrv "github.com/ant-libs-go/ant-mq/server/rpc"
)

var (
	addr         = flag.String("addr", ":9472", "MQ server listen address")
	pwd          = flag.String("d", "", "work directory")
	transport    = flag.String("transport", "rpc", "Transport for communication. Support rpc")
	pubWorkLimit = flag.Int("work", 10, "Max publish Goroutine count")

	// client flags
	interactive    = flag.Bool("i", false, "Interactive client mode")
	topic          = flag.String("topic", "", "Message topic name, multiple separated by commas")
	subscriber     = flag.String("subscriber", "", "Message subscriber, multiple separated by commas")
	register       = flag.Bool("register", false, "Register a new topic or subscriber")
	deregister     = flag.Bool("deregister", false, "Deregister a topic or subscriber")
	getTopics      = flag.Bool("gettopics", false, "Get topic list")
	getSubscribers = flag.Bool("getsubscribers", false, "Get subscriber list")

	// subscribe relation manager
	subscribe   = flag.Bool("subscribe", false, "Subscribe a new relation")
	unsubscribe = flag.Bool("unsubscribe", false, "Unsubscribe a relation")
)

func init() {
	flag.Parse()

	if *pwd == "" {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", "1.0.1")
	os.Setenv("WORKDIR", *pwd)

	if *interactive && (*register || *deregister) && (len(*topic) == 0 && len(*subscriber) == 0) {
		log.Fatal("Topic or subscriber not specified")
	}
	if *interactive && (*subscribe || *unsubscribe) && (len(*topic) == 0 || len(*subscriber) == 0) {
		log.Fatal("Topic and subscriber must specified")
	}
}

func main() {
	if *interactive {
		cli()
		return
	}

	core.DefaultBroker = core.NewBroker(
		core.WithTransport(*transport),
		core.WithPubGoroutineLimit(*pubWorkLimit),
	)

	err := core.LoadSubscribeRelation()
	if err != nil {
		log.Fatal(err)
	}

	options := []server.Option{
		server.WithListen(*addr),
	}
	var srv server.Server
	switch *transport {
	case "rpc":
		srv = rpcsrv.New(options...)
	default:
		log.Fatal("transport not support")
	}

	log.Printf("%s transport enabled", *transport)
	log.Printf("ANT-MQ listening on %s", *addr)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

// TODO
func cli() {
	switch {
	case *register:
	case *deregister:
	case *subscribe:
	case *unsubscribe:
	case *getTopics:
	case *getSubscribers:
	default:
		log.Fatal("The function of unrealized")
	}
}
