package main

import (
	"github.com/tomoyane/embedded-mysql-container/container"
)

func main() {
	containerDaemon := container.ContainerDaemonImpl{}.NewContainerDaemonImpl()

	containerDaemon.InitDocker()
	containerDaemon.PullImage("docker.io/library/mysql:5.7")

	containerId := containerDaemon.BuildImage(
		"mysql:5.7",
		"embedded_mysql3")

	containerDaemon.StartContainer(containerId)
	
	container.AddSchema("test")

	containerDaemon.SetupLogOfContainer(containerId)

	containerDaemon.StopContainer(containerId)
	containerDaemon.DeleteContainer(containerId)
}