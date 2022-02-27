DOCKER_IMAGE_NAME=go-service-rest-fiber-two
DOCKER_CONTAINER_NAME=go-fiber-two

start:
	go run main.go
run-examples-fib:
	cd examples/fib; go run fib.go
build:
	go build -ldflags="-s -w" -o apiserver .
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-run:
	docker run --rm -ti -d -p 3000:3000 --name ${DOCKER_CONTAINER_NAME} ${DOCKER_IMAGE_NAME}
	docker ps -a

docker-stop:
	docker stop ${DOCKER_CONTAINER_NAME}
	docker ps -a

air-install:
	# REF: https://github.com/cosmtrek/air
	sh air_install.sh
