DOCKER_IMAGE_NAME=go-fibonacci-service-fiber
DOCKER_CONTAINER_NAME=go-fibonacci-service-fiber-c
DOCKER_USERNAME=kevbaker

# HELP
help:
	@echo "\nTARGETS:\n"
	@make -qpRr | egrep -e '^[a-z].*:$$' | sed -e 's~:~~g' | sort
	@echo ""
list:
	make help

# TARGETS
start:
	go run main.go
start-reload:
	./bin/air
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
docker-kill:
	docker kill ${DOCKER_CONTAINER_NAME}

docker-push:
	docker image tag ${DOCKER_IMAGE_NAME}:latest ${DOCKER_USERNAME}/${DOCKER_IMAGE_NAME}:latest
	docker push ${DOCKER_USERNAME}/${DOCKER_IMAGE_NAME}:latest

air-install:
	sh ./extra/scripts/air_install.sh
