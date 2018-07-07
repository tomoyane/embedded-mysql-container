package main

type ContainerDaemon interface {
	InitDocker() bool

	PullImage(imageName string) bool

	BuildImage(imageName string, containerName string) string

	StartContainer(containerId string) bool

	StopContainer(containerId string) bool

	StopAllContainer()

	DeleteContainer(containerId string)

	WaitForContainer(containerId string)

	SetupLogOfContainer(containerId string)
}
