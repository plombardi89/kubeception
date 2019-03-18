ARG BASE_IMAGE="kubeception-base"
FROM ${BASE_IMAGE}

ARG KUBERNETES_VERSION="1.13.0"

COPY kubernetes.list /etc/apt/sources.list.d/kubernetes.list

ARG APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE="false"
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && apt-clean-install kubeadm kubectl kubelet

COPY bootstrap /usr/local/bin/bootstrap
RUN  chmod +x  /usr/local/bin/bootstrap

COPY bootstrap.service /etc/systemd/system
RUN  systemctl enable bootstrap.service
