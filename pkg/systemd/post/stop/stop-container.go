package stop

import (
	"flag"
	"os"

	"github.com/aledbf/systemd-go/pkg/docker"
	Log "github.com/aledbf/systemd-go/pkg/log"
)

var log = Log.New()

//  ExecStopPost=-/usr/bin/docker rm -f deis-builder
func main() {
	name := flag.String("name", "", "service to stop")

	flag.Parse()

	if flag.NFlag() < 1 {
		os.Exit(1)
	}

	if *name == "" {
		log.Fatal("invalid service name")
	}

	if err := docker.StopContainer(*name); err != nil {
		// Nothing (for now)
	}

	os.Exit(0)
}
