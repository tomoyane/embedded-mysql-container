package main

import (
	"embedded-mysql-container/daemon"
)

var container = daemon.ContainerDaemon{}

func main() {
	container.InitDocker()

	container.PullImage()

	var containerId = container.BuildImage()

	container.StartContainer(containerId)

	container.WaitForContainer(containerId)

	container.SetupLogOfContainer(containerId)
}