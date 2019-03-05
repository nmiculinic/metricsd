FROM golang:alpine AS build-env

WORKDIR /src
RUN apk add git
ADD . .
RUN go build -o metricsd ./cmd/metricsd

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/metricsd /app/
ENTRYPOINT ./metricsd
