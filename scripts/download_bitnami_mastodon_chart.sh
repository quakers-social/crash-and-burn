#!/usr/bin/env bash

set -x
set -e
set -u

export START_PWD=$(pwd)
export WORK_DIR="$(pwd)/pulumi"
SURCE_URL="https://github.com/bitnami/charts/archive/refs/heads/main.zip"

cd ${WORK_DIR}
mkdir -p ./charts

curl -L ${SURCE_URL} -o ./charts/charts-main.zip
unzip -u -d ./charts ./charts/charts-main.zip

ls -lah ./charts/charts-main/bitnami/
cd      ./charts/charts-main/bitnami/mastodon
helm dep update

