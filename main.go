package main

import (
	"github.com/tomoyane/embedded-mysql-container/container"
)

func main() {
	containerDaemon := container.ContainerDaemonImpl{}.New()
	embeddedMysql := container.MysqlConfigImpl{}.New()

	containerId := containerDaemon.StartEmbeddedMysql()
	embeddedMysql.AddSchema("test")

	containerDaemon.SetupLogOfContainer(containerId)
	containerDaemon.FinishEmbeddedMysql(containerId)
}
