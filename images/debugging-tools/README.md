DEBUGGING CONTAINER
=========================


BUILD
-----

```bash
podman build -t quakers-social/debugging:latest .
```

or without cache

```bash
podman build -t quakers-social/debugging:latest --no-cache=true .
```

IMAGE DEBUGGING
---------------

Start

```bash
podman run --rm -it  --entrypoint=sh quakers-social/debugging:latest
```


IMAGE PUSH BY HAND
------------------

```bash
podman login docker.io
LATES_VERSION=0.0.2
DOCKER_ACCOUNT=olafradicke
podman tag  quakers-social/debugging:latest  ${DOCKER_ACCOUNT}/debugging:${LATES_VERSION}
podman push ${DOCKER_ACCOUNT}/debugging:${LATES_VERSION}
```



NOTES
-----
