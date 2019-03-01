#!/bin/sh

cat << EOF > pack.Dockerfile
    FROM alpine
    RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
    WORKDIR /bin
    COPY metricsd .
    ENTRYPOINT [ "/bin/metricsd" ]
EOF
cat pack.Dockerfile

echo "building image tag ${IMAGE}"
docker build -f pack.Dockerfile -t ${IMAGE} .
docker push ${IMAGE}
echo pushed ${IMAGE}
