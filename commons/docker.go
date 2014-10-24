package commons

import (
	"github.com/fsouza/go-dockerclient"
)

const (
	endpoint = "unix:///var/run/docker.sock"
)

func StopContainer(name string) error {
	client, _ := docker.NewClient(endpoint)
	opts := docker.RemoveContainerOptions{ID: name, Force: true}
	return client.RemoveContainer(opts)
}
