# build stage
FROM golang:alpine AS build-env
WORKDIR /src
RUN apk update && apk add git gcc
ADD . .
RUN go build -v -x -o metricsd .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/metricsd /app/
ENTRYPOINT ./metricsd
