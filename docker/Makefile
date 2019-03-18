SHELL := /usr/bin/env bash

all: build.base build.node

build.base:
	docker build \
	-t kubeception:base \
	-t quay.io/plombardi89/kubeception:base \
	-f base/Dockerfile \
	base

build.node:
	cp ../bin/bootstrap node/bootstrap
	docker build \
	-t kubeception:node \
	-t quay.io/plombardi89/kubeception:node \
	-f node/Dockerfile \
	node

run.node:
	docker run \
	--rm \
	-it \
	--privileged \
	--name kubeception-node \
	-v /sys/fs/cgroup:/sys/fs/cgroup:ro \
	kubeception-node

exec.node:
	docker exec --env COLUMNS=$(shell tput cols) --env LINES=$(shell tput lines) -it kubeception-node /bin/bash

kill.node:
	docker kill kubeception-node

push:
	docker push quay.io/plombardi89/kubeception:base
	docker push quay.io/plombardi89/kubeception:node
