#!/usr/bin/env bash

set -x
set -e
set -u

go get github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes@latest || true
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
go mod tidy
mkdir ~/.kube
