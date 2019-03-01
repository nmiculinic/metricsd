all: build test

.PHONY: all
build:
	go build -o metricsd .
.PHONY: build
test:
	go test -race ./...
.PHONY: test
docker-build:
	echo "TODO"
.PHONY: docker-build
