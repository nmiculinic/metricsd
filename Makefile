IMAGE=gitlab.com/neven-miculinic/metricsd
TEST_DBURL=pg://postgres:root@localhost/postgres?sslmode=disable


GIT_SHA=$(shell git rev-parse HEAD)

all: build test

.PHONY: all
build:
	go build -o metricsd ./cmd/metricsd
.PHONY: build

static-build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -X github.com/nmiculinic/metricsd.Version=$(GIT_SHA)' -o metricsd ./cmd/metricsd

PHONY: build

run: build
	./metricsd --dburl $(TEST_DBURL)

test:
	go test -v -race ./...
.PHONY: test

local-test:
	TEST_DBURL=$(TEST_DBURL) go test -v -race ./...
.PHONY: test

docker: build
	IMAGE=$(IMAGE) ./.build_docker.sh
.PHONY: docker-build

proto: service.proto
	protoc -I .  --letmegrpc_out=pkg/metricsd --gogom_out=plugins=grpc:./pkg/metricsd service.proto

.PHONY: proto

clean:
	go clean -cache -modcache -testcache
.PHONY: clean

start-test-db:
	docker run --rm -it -p 5432:5432 --name postgresql -e POSTGRES_PASSWORD=root timescale/timescaledb:latest-pg11
.PHONY: start-test-db

psql-db:
	PGHOST=localhost PGPASSWORD=root PGUSER=postgres PGDATABASE=postgres psql
.PHONY: psql-db

bootstrap-db:
	PGHOST=localhost PGPASSWORD=root PGUSER=postgres PGDATABASE=postgres psql -f bootstrap.sql
.PHONY: bootsrap-db
