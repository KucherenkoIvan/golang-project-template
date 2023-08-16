FROM golang:latest as base

EXPOSE 4000

WORKDIR /app
RUN mkdir "/build"
COPY ./src .
COPY ./docker/build.sh /build.sh

FROM base as deamon

ENV CGO_ENABLED=0

RUN go get github.com/go-delve/delve/cmd/dlv@latest && \
    go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

RUN go get github.com/githubnemo/CompileDaemon && \
    go install github.com/githubnemo/CompileDaemon

FROM deamon

CMD CompileDaemon -build="sh /build.sh" -command="/go/bin/dlv --listen=:4000 --headless=true --api-version=2 --accept-multiclient exec /build/app" -color -log-prefix=false
