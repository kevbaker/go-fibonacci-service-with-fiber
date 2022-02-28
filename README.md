# Go Fibonacci Example Service with Fiber

## Overview

A simple service that outputs a fibonacci sequence with the number of items requested.

## Commands


### Start local service
```
make start
```

### Run local example fibonacci app
```
make run-examples-fib
```

### Build local app
```
make build
```

### Docker build container image
```
make docker-build
```

### Docker run container
```
make docker-run
```


### Docker stop container
```
make docker-stop
```

## Usage

* Run the service using the commands above
* Call service `curl -X GET http://localhost:3000/fibonacci?count=100`
* The above `count` set the numnber of items in the fibonacci sequence.
* NOTE: Larger `count` values will have a noticable impact on CPU.


## NOTES

-  [Live Reload](https://netboxify.com/blog/live-reload-development-environment-with-docker-and-go)
