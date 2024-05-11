FROM fedora

RUN dnf install -y \
    curl \
	telnet \
	traceroute \
	postgresql \
	mariadb \
	emacs-nox \
	bash \
	procps-ng

ENTRYPOINT ["top", "-b"]
#CENTRYPOINT ["/bin/bash/", "-c", "while true; do echo hello; sleep 10;done"]