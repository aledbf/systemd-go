This is a work in progress. 
===============

NOT READY TO USE!!!!
===============



Basically rewrite the current shell scripts in Go :)

This is just to improve the current output from the services and provide a way to customize the platform

from:
```console
[Service]
EnvironmentFile=/etc/environment
TimeoutStartSec=30m
ExecStartPre=/bin/sh -c "docker inspect deis-builder-data >/dev/null 2>&1 || docker run --name deis-builder-data -v /var/lib/docker ubuntu-debootstrap:14.04 /bin/true"
ExecStartPre=/bin/sh -c "IMAGE=`/run/deis/bin/get_image /deis/builder` && docker history $IMAGE >/dev/null || docker pull $IMAGE"
ExecStartPre=/bin/sh -c "docker inspect deis-builder >/dev/null && docker rm -f deis-builder || true"
ExecStart=/bin/sh -c "IMAGE=`/run/deis/bin/get_image /deis/builder` && docker run --name deis-builder --rm -p 2223:22 --volumes-from=deis-builder-data -e EXTERNAL_PORT=2223 -e HOST=$COREOS_PRIVATE_IPV4 --privileged $IMAGE"
ExecStartPost=/bin/sh -c "echo 'Waiting for builder on 2223/tcp...' && until echo 'dummy-value' | ncat $COREOS_PRIVATE_IPV4 2223 >/dev/null 2>&1; do sleep 1; done"
ExecStartPost=/bin/bash -c "nsenter --pid --uts --mount --ipc --net --target $(docker inspect --format='{{ .State.Pid }}' deis-builder) /usr/local/bin/push-images"
ExecStopPost=-/usr/bin/docker rm -f deis-builder
Restart=on-failure
RestartSec=5
```

to:
```console
[Service]
EnvironmentFile=/etc/environment
TimeoutStartSec=30m
ExecStartPre=/opt/bin/launch-data-container -name=deis-builder
ExecStartPre=/opt/bin/check-image -image=/deis/builder
ExecStartPre=/opt/bin/remove-running-container -name=deis-builder
ExecStart=/opt/bin/launch-container -image=/deis/builder -name=deis-builder
ExecStartPost=/opt/bin/wait-for-port -port=2223/tcp -text="Waiting for builder on 2223/tcp..."
ExecStartPost=/bin/bash -c "nsenter --pid --uts --mount --ipc --net --target $(docker inspect --format='{{ .State.Pid }}' deis-builder) /usr/local/bin/push-images"
ExecStopPost=/opt/bin/stop-container -name=deis-builder
Restart=on-failure
RestartSec=5
```

Every ExecStartPre line is just a small go app doing the same thing than before.
Why different apps? To keep it simple, 1 app -> 1 task

In ExecStart:
```console
ExecStart=/opt/bin/launch-container -image=/deis/builder -name=deis-builder
````

**Objectives:**

* Encapsulate the details
* Avoid shell mess
* use etcd to customize the service:
```
/deis/systemd/deis-builder/
                           variables
                           volumes
                           ports
```

doing this is possible to do other things:
```
/deis/systemd/app/defaults/
                           variables
                           volumes
/deis/systemd/app/example-php/
                           variables                                                                  
````

*In variables is possible to common custom variables without touching deis-controller python template*



* swallow the output.

  Show just what is necessary, things like `Unable to find image ubuntu-debootstrap:14.04 locally` are useless to new users.

* show more details

  Provide a way to get more details (something like in https://github.com/deis/deis/pull/2203) using /deis/debugMode to indicate this

* Add /deis/platform/useNodeNames (-h %H)

  Use the real node hostname inside the container (and not a generated one)


### TODO:

- [ ] check-image
- [ ] launch-data-container
- [ ] remove-running-container
- [ ] stop-container
- [x] wait-for-port
- [ ] start-component

