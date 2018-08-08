# Embedded MySQL container 
[![Build Status](http://www.concourse.developer-tm.com:8080/api/v1/teams/main/pipelines/embedded-mysql-container-pipeline/jobs/test/badge)](https://www.concourse.developer-tm.com/teams/main/pipelines/embedded-mysql-container-pipeline)

Embedded MySQL container.

This container's object want to resolve unit test database.

## Docker engine
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


### echo sample code

```

```

### gin sample code

```

```

### Revel sample code

```

```

