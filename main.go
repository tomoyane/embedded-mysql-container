package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"unittest-mysql-container/exception"
	"unittest-mysql-container/common"
)

var errorHandler = exception.ErrorHandler{}
var dockerContext = context.Background()
var dockerCli, cliErr = client.NewEnvClient()
var mysql string

func main() {
	if cliErr != nil {
		errorHandler.ErrorMessage(
			"docker engine is not working.",
		)
	}
}

func PullMysqlContainer(version string) {
	switch version {
	case common.VERSION56:
		mysql = "mysql:" + common.VERSION56

	case common.VERSION57:
		mysql = "mysql:" + common.VERSION57

	case common.VERSION_LATEST:
		mysql = "mysql"

	default:
		mysql = "mysql"
	}

	_, cliErr = dockerCli.ImagePull(dockerContext, mysql, types.ImagePullOptions{})
	if cliErr != nil {
		errorHandler.ErrorMessage(
			"docker pull command failed.",
		)
	}
}