IMAGE=gitlab.com/neven-miculinic/metricsd

all: build test

.PHONY: all
build:
	go build -o metricsd .
.PHONY: build
test:
	go test -race ./...
.PHONY: test

docker: build
	IMAGE=$(IMAGE) ./.build_docker.sh
.PHONY: docker-build
