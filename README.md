[![Build Status](http://www.concourse.developer-tm.com:8080/api/v1/teams/main/pipelines/embedded-mysql-container-pipeline/jobs/test/badge)](https://www.concourse.developer-tm.com/teams/main/pipelines/embedded-mysql-container-pipeline)

# Embedded MySQL container for Golang
Embedded MySQL container.

This container's object want to resolve unit test database.

## Docker setup
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

## How to use
Go get command.
```
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

