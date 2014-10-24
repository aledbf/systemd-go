FROM golang:1.3

WORKDIR /go/src/github.com/deis/deis/systemd

ADD . /go/src/github.com/deis/deis/systemd
RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/deis/deis/systemd
