SHELL := /usr/bin/env bash

GIT_COMMIT = $(shell git rev-parse --verify HEAD)

GOOS    = $(shell go env GOOS)
GOARCH  = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH) -ldflags "-X github.com/datawire/kubernaut/pkg/version.GitCommit=${GIT_COMMIT}"

BINARY_BASENAME = kubeception

all: clean fmt build

build: build.bootstrap build.kubeception

build.bootstrap:
	go build -o bin/bootstrap ./cmd/bootstrap/...

build.kubeception:
	$(GOBUILD) ./cmd/kubeception/...
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

clean:
	rm -rf bin

fmt:
	go fmt ./...
