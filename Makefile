
build:
	CGO_ENABLED=0 godep go build -a -ldflags '-s' post/start/wait-for-port.go

test:
	@echo no unit tests
