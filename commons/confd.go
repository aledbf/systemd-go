package commons

import (
	"os"
	"os/exec"
	"time"

	log "github.com/Sirupsen/logrus"
)

func WaitForInitialConfd(etcd string, timeout time.Duration) {
	for {
		cmd := exec.Command("confd", "-onetime", "-node", etcd, "-config-file", "/app/confd.toml")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err == nil {
			break
		}

		time.Sleep(timeout)
	}
}

func LaunchConfd(etcd string) {
	cmd := exec.Command("confd", "-node", etcd, "-config-file", "/app/confd.toml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal("confd terminated by error: %v", err)
	}
}
