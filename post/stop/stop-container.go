package stop

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/deis/systemd/commons"
)

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

	if err := commons.StopContainer(*name); err != nil {
		// Nothing (for now)
	}

	os.Exit(0)
}
