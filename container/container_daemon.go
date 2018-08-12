package container

type ContainerDaemon interface {
	InitDocker() bool

	PullImage(imageName string) bool

	BuildImage(imageName string, containerName string) string

	StartContainer(containerId string) bool

	StopContainer(containerId string) bool

	StopAllContainer()

	DeleteContainer(containerId string)

	SetupLogOfContainer(containerId string)

	WaitRun(containerId string)
}
