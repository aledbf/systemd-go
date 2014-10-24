package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/coreos/go-etcd/etcd"
	"github.com/deis/deis/systemd/commons"
	. "github.com/visionmedia/go-debug"
)

const (
	useNodeNames = "/deis/platform/useNodeNames"
)

var debug = Debug("systemd:start")

func main() {
	debug("starting service")

	host := commons.Getopt("COREOS_PRIVATE_IPV4", "127.0.0.1")
	etcdPort := commons.Getopt("ETCD_PORT", "4001")

	client := etcd.NewClient([]string{"http://" + host + ":" + etcdPort})

	// Wait for terminating signal
	exitChan := make(chan os.Signal, 2)
	signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)
	<-exitChan
}
