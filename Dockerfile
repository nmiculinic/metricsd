FROM golang:alpine AS build-env

WORKDIR /src
ADD . .
RUN go build -o metricsd .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/metricsd /app/
ENTRYPOINT ./metricsd
