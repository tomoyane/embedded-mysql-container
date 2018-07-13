package main

import (
	"context"
	"os"
	"io"
	"fmt"
	"time"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
)

type ContainerDaemonImpl struct {
	errorHandler *Error
}

var dockerContext = context.Background()
var dockerCli, cliErr = client.NewEnvClient()

func (c ContainerDaemonImpl) NewContainerDaemonImpl() ContainerDaemon {
	return &ContainerDaemonImpl{}
}

func (c ContainerDaemonImpl) InitDocker() bool {
	if cliErr != nil {
		c.errorHandler.ErrorMessage(
			"docker cli new failed.",
			cliErr,
		)

		return false
	}

	return true
}

func (c ContainerDaemonImpl) PullImage(imageName string) bool {
	reader, err := dockerCli.ImagePull(dockerContext, imageName, types.ImagePullOptions{})
	if err != nil {
		c.errorHandler.ErrorMessage(
			"docker pull image failed.",
			err,
		)

		return false

	}

	io.Copy(os.Stdout, reader)

	return true
}

func (c ContainerDaemonImpl) BuildImage(imageName string, containerName string) string {
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("3306"): []nat.PortBinding{nat.PortBinding{HostPort: "33306"}},
		},
	}

	nc := &network.NetworkingConfig{}

	resp, err := dockerCli.ContainerCreate(dockerContext, &container.Config{
		Image:        imageName,
		ExposedPorts: nat.PortSet{nat.Port("3306"): struct{}{}},
		Env:          []string{"MYSQL_ROOT_PASSWORD=root"},
		Tty:          true,
	},
		hc,
		nc,
		containerName,
	)

	if err != nil {
		c.errorHandler.ErrorMessage(
			"docker build failed.",
			err,
		)
	}

	return resp.ID
}

func (c ContainerDaemonImpl) StartContainer(containerId string) bool {
	if err := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); err != nil {
		c.errorHandler.ErrorMessage(
			"docker start failed.",
			err,
		)

		return false
	}

	return true
}

func (c ContainerDaemonImpl) StopContainer(containerId string) bool {
	timeout := 5 * time.Second
	if err := dockerCli.ContainerStop(dockerContext, containerId, &timeout); err != nil {
		c.errorHandler.ErrorMessage(
			"docker containerDaemon stop failed.",
			err,
		)

		return false
	}

	return true
}

func (c ContainerDaemonImpl) StopAllContainer() {
	containers, err := dockerCli.ContainerList(dockerContext, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, con := range containers {
		if err := dockerCli.ContainerStop(dockerContext, con.ID, nil); err != nil {
			c.errorHandler.ErrorMessage(
				"docker containerDaemon all stop failed.",
				err,
			)
		}
		fmt.Println("Success")
	}
}

func (c ContainerDaemonImpl) DeleteContainer(containerId string) {
	err := dockerCli.ContainerRemove(dockerContext, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		c.errorHandler.ErrorMessage(
			"docker container delete failed.",
			err,
		)
	}
}

func (c ContainerDaemonImpl) WaitForContainer(containerId string) {
	_, err := dockerCli.ContainerWait(dockerContext, containerId)
	if err != nil {
		c.errorHandler.ErrorMessage(
			"docker wait for start failed.",
			err,
		)
	}
}

func (c ContainerDaemonImpl) SetupLogOfContainer(containerId string) {
	out, err := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		c.errorHandler.ErrorMessage(
			"docker logging failed.",
			err,
		)
	}

	io.Copy(os.Stdout, out)
}
