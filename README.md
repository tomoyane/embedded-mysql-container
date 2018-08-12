# Embedded MySQL container 
[![Build Status](http://www.concourse.developer-tm.com:8080/api/v1/teams/main/pipelines/embedded-mysql-container-pipeline/jobs/test/badge)](https://www.concourse.developer-tm.com/teams/main/pipelines/embedded-mysql-container-pipeline)

Embedded MySQL container.

This container's object want to resolve unit test database.

## How to task ?
 - Your tests can run on production-like. 
 - You can easily try mysql linkage.

## Using Docker engine
Embedded MySQL container application needs to docker engine in your machine.

If you don't install docker engine, have to install docker engine.

[Docker engine](https://docs.docker.com/engine/)

## Architecture

```
 +----------- ApplicationTest ------------+  
 |                                        |
 |         +-------- 1 ---------+         |
 |         |  Develop machine   |         |
 |         |                    |         |
 |         |    docker pull     |         |
 |         |    docker build    |         |
 |         |    docker start    |         |
 |         |                    |         |
 |         +--------------------+         |
 |                   |                    |          
 |                   |                    |
 |         +-------- 2 ---------+         |  
 |         |      Unit Test     |         |
 |         |  Integration test  |         |
 |         +--------------------+         |
 |                                        |
 +----------------------------------------+  
```

### Default MySQL config
| Host | Port | User | Pass |
| --- | --- | --- | --- |
| 127.0.0.1 | 33306 | root | root |

## Usage
Go get command.
```
$ go get github.com/tomoyane/embedded-mysql-container
```

If you can get this error, you can do it after command.

This is Go vendor issue maybe.
```
cannot use exposedPorts (type map["github.com/docker/go-
connections/nat".Port]struct {}) as type map["github.com/docker/docker/vendor/github.com/docker/go-
connections/nat".Port]struct {} in field value
```

```
$ rm -rf $GOPATH/src/github.com/docker/docker/vendor/github.com/docker/go-connections
$ go get github.com/tomoyane/embedded-mysql-container
```

### Basic example

```go
import "github.com/tomoyane/embedded-mysql-container/container"

func main() {
    containerDaemon := container.ContainerDaemonImpl{}.New()
    embeddedMysql := container.MysqlConfigImpl{}.New()

    containerId := containerDaemon.StartEmbeddedMysql()

    // Need to sleep.
    // Wait for container startup.
    time.Sleep(10 * time.Second)

    embeddedMysql.AddSchema("test")

    // Init database
    
    containerDaemon.FinishEmbeddedMysql(containerId)
}
```

### Unit Test example

```go
import "github.com/tomoyane/embedded-mysql-container/container"

func TestMain(m *testing.M) {
	containerDaemon := container.ContainerDaemonImpl{}.New()
	embeddedMysql := container.MysqlConfigImpl{}.New()

	containerId := containerDaemon.StartEmbeddedMysql()

	time.Sleep(10 * time.Second)

	embeddedMysql.AddSchema("auth_server")
	embeddedMysql.CreateTable("CREATE TABLE test.items (" +
		"id int(11) NOT NULL AUTO_INCREMENT," +
		"name varchar(128) NOT NULL," +
		"created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (id)," +
		"ENGINE=InnoDB DEFAULT CHARSET=utf8")
	
	// Init main database
	
	code := m.Run()

	// After finish mysql container
	containerDaemon.FinishEmbeddedMysql(containerId)

	os.Exit(code)
}
```

### Custom example
If you want set any mysql version, you can set version mysql.

And If you want set any container name, you can set container name. 

```go
import "github.com/tomoyane/embedded-mysql-container/container"

func main() {
    containerDaemon := container.ContainerDaemonImpl{}.New()
    embeddedMysql := container.MysqlConfigImpl{}.New()
    
    containerDaemon.InitDocker()
    
    // MySQL5.7
    containerDaemon.PullImage("docker.io/library/mysql:5.7")
    containerId := containerDaemon.BuildImage(
        "mysql:5.7",
        "embedded_mysql_container")
    containerDaemon.StartContainer(containerId)
}
```