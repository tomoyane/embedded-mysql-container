package daemon

import (
	"context"
	"os"
	"io"
	"fmt"
	"embedded-mysql-container/exception"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/network"
	"time"
)

type ContainerDaemon struct {
}

var dockerContext = context.Background()
var dockerCli, cliErr = client.NewEnvClient()
var errorHandler = exception.ErrorHandler{}

func (c ContainerDaemon) InitDocker() {
	if cliErr != nil {
		errorHandler.ErrorMessage(
			"docker cli new failed.",
			cliErr,
		)
	}
}

func (c ContainerDaemon) PullImage() {
	reader, pullErr := dockerCli.ImagePull(dockerContext, "docker.io/library/mysql:5.7", types.ImagePullOptions{})
	if pullErr != nil {
		errorHandler.ErrorMessage(
			"docker pull image failed.",
			pullErr,
		)
	}

	io.Copy(os.Stdout, reader)
}

func (c ContainerDaemon) BuildImage() string {
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("3306"): []nat.PortBinding{nat.PortBinding{HostPort: "33306"}},
		},
	}

	nc := &network.NetworkingConfig{}

	containerName := "embedded_mysql"

	resp, buildErr := dockerCli.ContainerCreate(dockerContext, &container.Config{
		Image:        "mysql:5.7",
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

func (c ContainerDaemon) StartContainer(containerId string) {
	if startErr := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); startErr != nil {
		errorHandler.ErrorMessage(
			"docker start failed.",
			startErr,
		)
	}
}

func (c ContainerDaemon) StopContainer(containerId string) {
	timeout := 5 * time.Second
	if stopErr := dockerCli.ContainerStop(dockerContext, containerId, &timeout); stopErr != nil {
		errorHandler.ErrorMessage(
			"docker container stop failed.",
			stopErr,
		)
	}
}

func (c ContainerDaemon) StopAllContainer() {
	containers, err := dockerCli.ContainerList(dockerContext, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, con := range containers {
		if err := dockerCli.ContainerStop(dockerContext, con.ID, nil); err != nil {
			errorHandler.ErrorMessage(
				"docker container all stop failed.",
				err,
			)
		}
		fmt.Println("Success")
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
	out, logErr := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if logErr != nil {
		errorHandler.ErrorMessage(
			"docker logging failed.",
			logErr,
		)
	}

	io.Copy(os.Stdout, out)
}
