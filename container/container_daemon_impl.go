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
	errContainer ErrorContainer
}

var (
	dockerContext = context.Background()
	dockerCli, cliErr = client.NewEnvClient()
)

func (c ContainerDaemonImpl) New() ContainerDaemon {
	return &ContainerDaemonImpl{}
}

func (c ContainerDaemonImpl) InitDocker() bool {
	if cliErr != nil {
		c.errContainer.msg = "docker cli new failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()

		return false
	}

	return true
}

func (c ContainerDaemonImpl) PullImage(imageName string) bool {
	reader, err := dockerCli.ImagePull(dockerContext, imageName, types.ImagePullOptions{})
	if err != nil {
		c.errContainer.msg = "docker pull image failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()

		return false

	}

	io.Copy(os.Stdout, reader)

	return true
}

func (c ContainerDaemonImpl) BuildImage(imageName string, containerName string) string {
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("3306"): []nat.PortBinding{{HostPort: "33306"}},
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
		c.errContainer.msg = "docker build failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()
	}

	return resp.ID
}

func (c ContainerDaemonImpl) StartContainer(containerId string) bool {
	if err := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); err != nil {
		c.errContainer.msg = "docker start failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()

		return false
	}

	return true
}

func (c ContainerDaemonImpl) StopContainer(containerId string) bool {
	timeout := 5 * time.Second
	if err := dockerCli.ContainerStop(dockerContext, containerId, &timeout); err != nil {
		c.errContainer.msg = "docker containerDaemon stop failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()

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
			c.errContainer.msg = "docker containerDaemon all stop failed."
			c.errContainer.error = cliErr
			c.errContainer.ErrorMessage()
		}
		fmt.Println("Success")
	}
}

func (c ContainerDaemonImpl) DeleteContainer(containerId string) {
	err := dockerCli.ContainerRemove(dockerContext, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		c.errContainer.msg = "docker container delete failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()
	}
}

func (c ContainerDaemonImpl) SetupLogOfContainer(containerId string) {
	out, err := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		c.errContainer.msg = "docker logging failed."
		c.errContainer.error = cliErr
		c.errContainer.ErrorMessage()
	}

	io.Copy(os.Stdout, out)
}

func (c ContainerDaemonImpl) WaitRun(containerId string) {
	dockerCli.ContainerWait(dockerContext, containerId)
}

func (c ContainerDaemonImpl) StartEmbeddedMysql() string {
	c.InitDocker()
	c.PullImage("docker.io/library/mysql:5.7")

	containerId := c.BuildImage(
		"mysql:5.7",
		"embedded_mysql3")

	c.StartContainer(containerId)

	time.Sleep(10 * time.Second)

	return containerId
}

func (c ContainerDaemonImpl) FinishEmbeddedMysql(containerId string) {
	c.StopContainer(containerId)
	c.DeleteContainer(containerId)
}