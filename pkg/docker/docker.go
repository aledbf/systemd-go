package docker

import (
	"github.com/aledbf/systemd-go/pkg/os"

	"github.com/fsouza/go-dockerclient"
)

const (
	defaultSocket = "unix:///var/run/docker.sock"
)

// NewDockerClient returns a new docker client
// By default the socket /var/run/docker.sock is used if there is no env DOCKER_HOST
func NewDockerClient() (*docker.Client, error) {
	endpoint := os.Getopt("DOCKER_HOST", defaultSocket)
	return docker.NewClient(endpoint)
}

func StopContainer(name string) error {
	client, _ := NewDockerClient()
	opts := docker.RemoveContainerOptions{ID: name, Force: true}
	return client.RemoveContainer(opts)
}
