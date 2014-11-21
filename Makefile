
BUILD_IMAGE := deis/systemd-build

build:
	docker build -t $(BUILD_IMAGE) .
	#docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/check-image bin
	#docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/launch-data-container bin
	#docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/remove-running-container bin
	#docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/stop-container bin
	docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/wait-for-port bin
	#docker cp `docker run -d $(BUILD_IMAGE)`:/go/bin/start-component bin

clean: check-docker check-registry
	docker rmi $(RELEASE_IMAGE) $(REMOTE_IMAGE)

test:
	@echo no unit tests
