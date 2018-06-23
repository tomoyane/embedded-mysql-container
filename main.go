package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"os"
	"io"
	"embedded-mysql-container/exception"
)

var dockerContext = context.Background()
var dockerCli, cliErr = client.NewEnvClient()
var errorHandler = exception.ErrorHandler{}

func main() {
	InitDocker()

	PullImage()

	var containerId = BuildImage()
	
	StartContainer(containerId)

	WaitForContainer(containerId)

	SetupLogOfContainer(containerId)
}

func InitDocker() {
	if cliErr != nil {
		errorHandler.ErrorMessage(
			"docker cli new failed.",
			cliErr,
		)
	}
}

func PullImage() {
	reader, pullErr := dockerCli.ImagePull(dockerContext, "docker.io/library/mysql", types.ImagePullOptions{})
	if pullErr != nil {
		errorHandler.ErrorMessage(
			"docker pull image failed.",
			pullErr,
		)
	}
	io.Copy(os.Stdout, reader)
}

func BuildImage() string {
	resp, buildErr := dockerCli.ContainerCreate(dockerContext, &container.Config {
		Image: "mysql",
		Cmd:   []string{"echo", "hello world"},
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

func StartContainer(containerId string) {
	if startErr := dockerCli.ContainerStart(dockerContext, containerId, types.ContainerStartOptions{}); startErr != nil {
		errorHandler.ErrorMessage(
			"docker start failed.",
			startErr,
		)
	}
}

func WaitForContainer(containerId string)  {
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

func SetupLogOfContainer(containerId string) {
	out, logErr := dockerCli.ContainerLogs(dockerContext, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if logErr != nil {
		errorHandler.ErrorMessage(
			"docker logging failed.",
			logErr,
		)
	}

	io.Copy(os.Stdout, out)
}