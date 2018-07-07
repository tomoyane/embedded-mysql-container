package main

func main() {
	containerDaemon := ContainerDaemonImpl{}.NewContainerDaemonImpl()

	containerDaemon.InitDocker()
	containerDaemon.PullImage("docker.io/library/mysql:5.7")

	containerId := containerDaemon.BuildImage(
		"mysql:5.7",
		"embedded_mysql",
	)

	containerDaemon.StartContainer(containerId)
	containerDaemon.SetupLogOfContainer(containerId)
	containerDaemon.StopContainer(containerId)
}