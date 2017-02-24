# Infinite logger

A stoopid-simple app that logs a message every 5 seconds.

## Uh, why?

Sometimes you want a really light-weight binary to test deployments.

## Installation

### Native Binary

Assuming a [correctly configured](https://golang.org/doc/install#testing)
Go toolchain:

```
go get github.com/cpliakas/infinite-logger
```

### Docker Image

Assuming a working [Docker installation](https://docs.docker.com/engine/installation/):

```
docker run cpliakas/infinite-logger
```

## Building the Docker Image

Basically this is copy-pasta for me ...

```
GOOS=linux go build .
docker build -t cpliakas/infinite-logger .
```
