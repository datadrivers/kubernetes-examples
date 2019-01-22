#
# Base build image
FROM golang:1.11-alpine AS build_base

# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git
WORKDIR /go/src/github.com/datadrivers/kubernetes-examples/webserver

# Force the go compiler to use modules and copy the module
# definition files into source directory
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

#
# Intermediate build image
FROM build_base AS builder
COPY . .
RUN GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"'

#
# Runtime image
FROM alpine AS simple_webserver
LABEL maintainer Marcus Franke <marcus.franke@datadrivers.de>

RUN apk add ca-certificates

# Copy the binary from the builder image into the runtime image
COPY --from=builder /go/bin/kubernetes-examples /bin/simple_webserver
ENTRYPOINT ["/bin/simple_webserver"]