#
# Base build image
FROM golang:1.11-alpine AS build_base

# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git
WORKDIR /go/src/github.com/datadrivers

# Force the go compiler to use modules and copy the module
# definition files into source directory

#
# Intermediate build image
FROM build_base AS builder

ENV GOOS=linux
ENV GOARCH=amd64
RUN go get github.com/rakyll/statik/fs
RUN git clone https://github.com/datadrivers/kubernetes-examples
WORKDIR /go/src/github.com/datadrivers/kubernetes-examples/webserver
RUN go build -o /simple_webserver

#
# Runtime image
FROM alpine AS simple_webserver
LABEL maintainer Marcus Franke <marcus.franke@datadrivers.de>

RUN apk add ca-certificates

# Copy the binary from the builder image into the runtime image
COPY --from=builder simple_webserver /bin/simple_webserver
ENTRYPOINT ["/bin/simple_webserver"]
