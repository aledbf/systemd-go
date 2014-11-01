package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/go-etcd/etcd"
	"github.com/deis/systemd/commons"
	"github.com/deis/systemd/logger"
)

const (
	useNodeNames = "/deis/platform/useNodeNames"
)

func main() {
	logger.Log.Debugf("starting service")

	host := commons.Getopt("COREOS_PRIVATE_IPV4", "127.0.0.1")
	etcdPort := commons.Getopt("ETCD_PORT", "4001")

	client := etcd.NewClient([]string{"http://" + host + ":" + etcdPort})
	client.SetConsistency(etcd.STRONG_CONSISTENCY)

	// Wait for terminating signal
	exitChan := make(chan os.Signal, 2)
	signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)
	<-exitChan
}
