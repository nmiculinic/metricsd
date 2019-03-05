IMAGE=gitlab.com/neven-miculinic/metricsd

all: build test

.PHONY: all
build:
	go build -o metricsd .
.PHONY: build

run: build
	./metricsd

test:
	go test -race ./...
.PHONY: test

docker: build
	IMAGE=$(IMAGE) ./.build_docker.sh
.PHONY: docker-build

proto: service.proto
	protoc -I . --gogom_out=plugins=grpc:./pkg/metricsd service.proto

.PHONY: proto
