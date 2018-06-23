package daemon

import (
	"embedded-mysql-container/exception"
	"context"
	"os"
	"io"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
	reader, pullErr := dockerCli.ImagePull(dockerContext, "docker.io/library/mysql", types.ImagePullOptions{})
	if pullErr != nil {
		errorHandler.ErrorMessage(
			"docker pull image failed.",
			pullErr,
		)
	}

	io.Copy(os.Stdout, reader)
}

func (c ContainerDaemon) BuildImage() string {
	resp, buildErr := dockerCli.ContainerCreate(dockerContext, &container.Config{
		Image: "mysql",
		Cmd:   []string{"-p", "3306:3306"},
		Tty:   true,
	},
		nil,
		nil,
		"",
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

func (c ContainerDaemon) StopAllContainer() {
	containers, err := dockerCli.ContainerList(dockerContext, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, con := range containers {
		if err := dockerCli.ContainerStop(dockerContext, con.ID, nil); err != nil {
			errorHandler.ErrorMessage(
				"docker stop failed.",
				err,
			)
		}
		fmt.Println("Success")
	}
}

func (c ContainerDaemon) WaitForContainer(containerId string) {
	statusCh, errCh := dockerCli.ContainerWait(dockerContext, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			errorHandler.ErrorMessage(
				"docker wait for start failed.",
				err,
			)
		}
	case <-statusCh:
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
