ARG BASE_IMAGE="ubuntu:18.04"
FROM ${BASE_IMAGE}

ENV LC_ALL C.UTF-8

# Ensure systemd is aware it is in a container
# Ref: https://www.freedesktop.org/wiki/Software/systemd/ContainerInterface/
ENV container docker

# Ensure systemd exits on SIGRTMIN+3, not SIGTERM (which re-executes it)
# Ref: https://bugzilla.redhat.com/show_bug.cgi?id=1201657
STOPSIGNAL SIGRTMIN+3

# Set DEBIAN_FRONTEND to noninteractive in order to suppress some apt warnings
ARG DEBIAN_FRONTEND=noninteractive

# Set Docker version
ARG DOCKER_VERSION="18.06.*"

# Set CNI version
ARG CNI_ARCH="amd64"
ARG CNI_VERSION="0.6.0"
ARG CNI_BASE_URL="https://storage.googleapis.com/kubernetes-release/network-plugins/"
ARG CNI_TARBALL="cni-plugins-${CNI_ARCH}-v${CNI_VERSION}.tgz"

COPY apt-clean-install /usr/local/bin/apt-clean-install
RUN  chmod +x /usr/local/bin/apt-clean-install

# Get dependencies
#
# The order of the apt-get install is:
#   1. packages necessary for docker
#   2. packages necessary for systemd
#   3. packages necessary for kubernetes
#   4. miscellaneous dev tools
#
# Afterwards we purge a bunch of leftover data from the filesystem then we delete a number of unnecessary
# Systemd services and disable KMsg in Journald.
RUN apt-clean-install \
        apt-transport-https ca-certificates curl software-properties-common gnupg2 lsb-release \
        systemd systemd-sysv libsystemd0 \
        conntrack iptables iproute2 ethtool socat util-linux mount ebtables udev kmod aufs-tools \
        bash \
    && find /lib/systemd/system/sysinit.target.wants/ -name "systemd-tmpfiles-setup.service" -delete \
    && rm -f /lib/systemd/system/multi-user.target.wants/* \
    && rm -f /etc/systemd/system/*.wants/* \
    && rm -f /lib/systemd/system/local-fs.target.wants/* \
    && rm -f /lib/systemd/system/sockets.target.wants/*udev* \
    && rm -f /lib/systemd/system/sockets.target.wants/*initctl* \
    && rm -f /lib/systemd/system/basic.target.wants/* \
    && echo "ReadKMsg=no" >> /etc/systemd/journald.conf

# Install Docker and CNI
# Silence "Warning: apt-key output should not be parsed (stdout is not a terminal)"
ARG APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE="false"
RUN curl -fsSL "https://download.docker.com/linux/ubuntu/gpg" | apt-key add - \
    && apt-key fingerprint 0EBFCD88 \
    && add-apt-repository \
        "deb https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
    && apt-clean-install "docker-ce=${DOCKER_VERSION}" \
    && curl -sSL --retry 5 --output /tmp/cni.tgz "${CNI_BASE_URL}${CNI_TARBALL}" \
    && mkdir -p /opt/cni/bin \
    && tar -C /opt/cni/bin -xzf /tmp/cni.tgz \
    && rm -rf /tmp/cni.tgz

ENTRYPOINT ["/sbin/init"]
VOLUME     ["/var/lib/docker"]