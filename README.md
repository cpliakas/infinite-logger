# Infinite logger

A stoopid-simple app that logs a message every 5 seconds.

## Uh, why?

Sometimes you want a really light-weight binary to test deployments.

## Docker build

Assuming a [correctly configured](https://golang.org/doc/install#testing)
Go toolchain:

```
GOOS=linux go build .
```

Assuming a working [Docker installation](https://docs.docker.com/engine/installation/):

```
docker build -t cpliakas/infinite-logger .
```
