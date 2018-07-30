package container

import (
	"io"
	"os"
	"time"
	"fmt"
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/network"
)

type ContainerDaemonImpl struct {
}

var (
	dockerContext = context.Background()
	dockerCli, cliErr = client.NewEnvClient()
)

func (c ContainerDaemonImpl) NewContainerDaemonImpl() ContainerDaemon {
	return &ContainerDaemonImpl{}
}

func (c ContainerDaemonImpl) InitDocker() bool {
	if cliErr != nil {
		errContainer := ErrorContainer{
			msg: "docker cli new failed.",
			error: cliErr,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (c ContainerDaemonImpl) PullImage(imageName string) bool {
	reader, err := dockerCli.ImagePull(dockerContext, imageName, types.ImagePullOptions{})
	if err != nil {
		errContainer := ErrorContainer{
			msg: "docker pull image failed.",
			error: err,
		}

		errContainer.ErrorMessage()

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
		errContainer := ErrorContainer{
			msg: "docker build failed.",
			error: err,
		}

		errContainer.ErrorMessage()
	}

	return resp.ID
}

func (c ContainerDaemonImpl) StartContainer(containerId string) bool {
	if err := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); err != nil {
		errContainer := ErrorContainer{
			msg: "docker start failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (c ContainerDaemonImpl) StopContainer(containerId string) bool {
	timeout := 5 * time.Second
	if err := dockerCli.ContainerStop(dockerContext, containerId, &timeout); err != nil {
		errContainer := ErrorContainer{
			msg: "docker containerDaemon stop failed.",
			error: err,
		}

		errContainer.ErrorMessage()

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
			errContainer := ErrorContainer{
				msg: "docker containerDaemon all stop failed.",
				error: err,
			}

			errContainer.ErrorMessage()
		}
		fmt.Println("Success")
	}
}

func (c ContainerDaemonImpl) DeleteContainer(containerId string) {
	err := dockerCli.ContainerRemove(dockerContext, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		errContainer := ErrorContainer{
			msg: "docker container delete failed.",
			error: err,
		}

		errContainer.ErrorMessage()
	}
}

func (c ContainerDaemonImpl) WaitForContainer(containerId string) {
	_, err := dockerCli.ContainerWait(dockerContext, containerId)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "docker wait for start failed.",
			error: err,
		}

		errContainer.ErrorMessage()
	}
}

func (c ContainerDaemonImpl) SetupLogOfContainer(containerId string) {
	out, err := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		errContainer := ErrorContainer{
			msg: "docker logging failed.",
			error: err,
		}

		errContainer.ErrorMessage()
	}

	io.Copy(os.Stdout, out)
}
