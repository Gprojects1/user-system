COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

export GO111MODULE=on

.PHONY: docker
docker: 
	@echo "-- building docker container - multistage"
	docker build -f build/Dockerfile.Multistage -t user-system .

.PHONY: docker_run
docker_run: 
	@echo "-- starting docker container"
	docker run -it -p 8080:8080 user-system

.PHONY: dc
dc: 
	@echo "-- starting docker compose"
	docker-compose -f ./deployments/docker-compose.yml up

.PHONY: dcb
dcb: 
	@echo "-- starting docker compose with build"
	docker-compose -f ./deployments/docker-compose.yml up --build
