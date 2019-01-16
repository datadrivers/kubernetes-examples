FROM golang:latest as build

RUN go get github.com/datadrivers/kubernetes-examples/webserver
WORKDIR /go/src/github.com/datadrivers/kubernetes-examples/webserver
RUN go build -o simple_webserver


FROM centos:7
LABEL maintainer Marcus Franke <marcus.franke@datadrivers.de>

COPY --from=build /go/src/github.com/datadrivers/kubernetes-examples/webserver/simple_webserver /bin

EXPOSE 8080
CMD ["/bin/simple_webserver"]
