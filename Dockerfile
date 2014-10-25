FROM golang:1.3

RUN go get github.com/tools/godep

WORKDIR /go/src/github.com/deis/systemd

ADD . /go/src/github.com/deis/systemd

RUN CGO_ENABLED=0 godep go build -a -ldflags '-s' -o github.com/deis/systemd .
