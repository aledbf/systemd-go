package start

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//  ExecStartPost=/opt/bin/wait-for-port -port=2223/tcp -text="Waiting for builder on 2223/tcp..."
func main() {
	port := flag.Int("port", 0, "Name of the image to check or download")
	text := flag.String("text", "waiting for port", "Text to show while the port is not available")
	timeout := flag.Int("timeout", 0, "Max wait time for the port in seconds. By default wait forever")

	flag.Parse()

	if flag.NFlag() < 3 {
		os.Exit(1)
	}

	if *port == 0 {
		log.Fatal("invalid port")
	}

	if *timeout < 0 {
		log.Fatal("invalid timeout")
	}

	timeoutInSeconds := (time.Duration(*timeout) * time.Second).Seconds()

	// start measuring the time to know how long it took
	startTime := time.Now()

	for {
		if len(*text) > 0 {
			log.Info(text)
		}

		if _, err := net.DialTimeout("tcp", "127.0.0.1:"+string(*port), 1); err == nil {
			break
		}

		// check if there is a max timeout
		if *timeout > 0 {
			elapsed := time.Since(startTime)
			if elapsed.Seconds() > timeoutInSeconds {
				log.Warn(os.Stderr, "timeout reached. aborting...")
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}
