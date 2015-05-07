package main

import (
	"flag"
	"math"
	"net"
	"os"
	"strconv"
	"time"

	Log "github.com/aledbf/systemd-go/pkg/log"
)

var log = Log.New()

//  ExecStartPost=/opt/bin/wait-for-port -port=2223 -text="Waiting for builder on 2223/tcp..."
func main() {
	ip := flag.String("ip", "0.0.0.0", "Ip address")
	port := flag.Int("port", 0, "Name of the image to check or download")
	text := flag.String("text", "waiting for port", "Text to show while the port is not available")
	// protocol := flag.String("text", "tcp", "Text to show while the port is not available")
	// response := flag.String("text", "", "Text to search in the response of the connection (only for body in protocol http)")
	timeout := flag.Int("timeout", 0, "Max wait time for the port in seconds. By default wait forever")

	flag.Parse()

	if flag.NFlag() < 1 {
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

	ipPort := *ip + ":" + strconv.Itoa(*port)

	log.Debugf("Waiting for %v to be available ", ipPort)

	for {
		if len(*text) > 0 {
			log.Info(*text)
		}

		if _, err := net.DialTimeout("tcp", ipPort, 1*time.Second); err == nil {
			elapsed := time.Since(startTime)
			log.Infof("TCP connection to %v available after %v seconds", ipPort,
				math.Ceil(elapsed.Seconds()))
			break
		}

		time.Sleep(1 * time.Second)

		// check if there is a max timeout
		if *timeout > 0 {
			elapsed := time.Since(startTime)
			if elapsed.Seconds() > timeoutInSeconds {
				log.Error("timeout reached. aborting...")
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}
