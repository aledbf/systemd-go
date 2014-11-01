package start

import (
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/deis/systemd/logger"
)

func main() {
	name := flag.String("name", "", "Name of the image to use")

	flag.Parse()

	if flag.NFlag() < 1 {
		os.Exit(1)
	}

	if !checkRunningContainer(name) {
		startDataContainer(name)
	}

	if *name == "" {
		logger.Log.Fatal("invalid image name")
	}

	os.Exit(0)
}

func checkRunningContainer(containerName string) bool {
	logger.Log.Debugf("checking if there is a container with name %s is running", containerName)
	cmd := exec.Command("docker", "inspect", containerName)
	if err := cmd.Run(); err != nil {
		return false
	} else {
		return true
	}
}

func startDataContainer(containerName string) bool {
	logger.Log.Debugf("launching data container name %s", containerName)
	cmd := exec.Command("docker", "run", "--name", containerName, "-v", "/var/lib/docker", "ubuntu-debootstrap:14.04", "/bin/true")
	if err := cmd.Run(); err != nil {
		return false
	} else {
		return true
	}
}
