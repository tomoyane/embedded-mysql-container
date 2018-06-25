package main

var containerDaemon = ContainerDaemon{}

func main() {
	containerDaemon.InitDocker()

	containerDaemon.PullImage("docker.io/library/mysql:5.7")

	var containerId = containerDaemon.BuildImage(
		"mysql:5.7",
		"embedded_mysql",
	)

	containerDaemon.StartContainer(containerId)

	containerDaemon.WaitForContainer(containerId)

	containerDaemon.SetupLogOfContainer(containerId)
}