package test

import (
	"testing"
	"github.com/tomoyane/embedded-mysql-container/container"
)

var containerDaemon = container.ContainerDaemonImpl{}.NewContainerDaemonImpl()

func TestInitContainer(t *testing.T) {
	result := containerDaemon.InitDocker()
	if !result {
		t.Failed()
	}
}

func TestPullImage(t *testing.T) {
	result := containerDaemon.PullImage("docker.io/library/alpine")

	if !result {
		t.Failed()
	}
}

func TestBuildImage(t *testing.T) {
	var containerId = containerDaemon.BuildImage("alpine", "alpine")
	if containerId == "" {
		t.Failed()
	}

	containerDaemon.StopContainer(containerId)
	containerDaemon.DeleteContainer(containerId)
}
