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

type ContainerDaemon struct {
}

var dockerContext = context.Background()
var dockerCli, cliErr = client.NewEnvClient()
var errorHandler = Error{}

func (c ContainerDaemon) InitDocker() bool {
	if cliErr != nil {
		errorHandler.ErrorMessage(
			"docker cli new failed.",
			cliErr,
		)

		return false
	}

	return true
}

func (c ContainerDaemon) PullImage(imageName string) bool {
	reader, pullErr := dockerCli.ImagePull(dockerContext, imageName, types.ImagePullOptions{})
	if pullErr != nil {
		errorHandler.ErrorMessage(
			"docker pull image failed.",
			pullErr,
		)

		return false

	}

	io.Copy(os.Stdout, reader)

	return true
}

func (c ContainerDaemon) BuildImage(imageName string, containerName string) string {
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("3306"): []nat.PortBinding{nat.PortBinding{HostPort: "33306"}},
		},
	}

	nc := &network.NetworkingConfig{}

	resp, buildErr := dockerCli.ContainerCreate(dockerContext, &container.Config{
		Image:        imageName,
		ExposedPorts: nat.PortSet{nat.Port("3306"): struct{}{}},
		Env:          []string{"MYSQL_ROOT_PASSWORD=root"},
		Tty:          true,
	},
		hc,
		nc,
		containerName,
	)

	if buildErr != nil {
		errorHandler.ErrorMessage(
			"docker build failed.",
			buildErr,
		)
	}

	return resp.ID
}

func (c ContainerDaemon) StartContainer(containerId string) bool {
	if startErr := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); startErr != nil {
		errorHandler.ErrorMessage(
			"docker start failed.",
			startErr,
		)

		return false
	}

	return true
}

func (c ContainerDaemon) StopContainer(containerId string) bool {
	timeout := 5 * time.Second
	if stopErr := dockerCli.ContainerStop(dockerContext, containerId, &timeout); stopErr != nil {
		errorHandler.ErrorMessage(
			"docker containerDaemon stop failed.",
			stopErr,
		)

		return false
	}

	return true
}

func (c ContainerDaemon) StopAllContainer() {
	containers, err := dockerCli.ContainerList(dockerContext, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, con := range containers {
		if err := dockerCli.ContainerStop(dockerContext, con.ID, nil); err != nil {
			errorHandler.ErrorMessage(
				"docker containerDaemon all stop failed.",
				err,
			)
		}
		fmt.Println("Success")
	}
}

func (c ContainerDaemon) DeleteContainer(containerId string) {
	err := dockerCli.ContainerRemove(dockerContext, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		errorHandler.ErrorMessage(
			"docker container delete failed.",
			err,
		)
	}
}

func (c ContainerDaemon) WaitForContainer(containerId string) {
	_, errCh := dockerCli.ContainerWait(dockerContext, containerId)
	if errCh != nil {
		errorHandler.ErrorMessage(
			"docker wait for start failed.",
			errCh,
		)
	}
}

func (c ContainerDaemon) SetupLogOfContainer(containerId string) {
	out, errLog := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if errLog != nil {
		errorHandler.ErrorMessage(
			"docker logging failed.",
			errLog,
		)
	}

	io.Copy(os.Stdout, out)
}
