
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o bin/wait-for-port pkg/systemd/post/start/wait-for-port.go

test:
	@echo no unit tests
