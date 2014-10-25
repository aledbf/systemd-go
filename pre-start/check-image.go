package preStart

import (
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

//  ExecStartPre=/bin/sh -c "IMAGE=`/run/deis/bin/get_image /deis/builder` && docker history $IMAGE >/dev/null || docker pull $IMAGE"
//  ExecStartPre=/opt/bin/check-image -image=/deis/builder
func main() {
	image := flag.String("image", "", "Name of the image to check or download")

	flag.Parse()

	if flag.NFlag() < 1 {
		os.Exit(1)
	}

	if !checkRunningContainer(name) {
		startDataContainer(name)
	}

	if *image == "" {
		log.Fatal("invalid image name")
	}

	os.Exit(0)
}

func checkRunningContainer(containerName string) bool {
	log.Debugf("checking if there is a container with name %s is running", containerName)
	cmd := exec.Command("docker", "inspect", containerName)
	if err := cmd.Run(); err != nil {
		return false
	} else {
		return true
	}
}

func startDataContainer(containerName string) bool {
	log.Debugf("launching data container name %s", containerName)
	cmd := exec.Command("docker", "run", "--name", containerName, "-v", "/var/lib/docker", "ubuntu-debootstrap:14.04", "/bin/true")
	if err := cmd.Run(); err != nil {
		return false
	} else {
		return true
	}
}
