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
}

//containerDaemon.InitDocker()
//containerDaemon.PullImage("docker.io/library/mysql:5.7")
//containerId := containerDaemon.BuildImage(
//"mysql:5.7",
//"embedded_mysql_container")
//containerDaemon.StartContainer(containerId)
