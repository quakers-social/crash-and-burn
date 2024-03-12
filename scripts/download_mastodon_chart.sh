#!/usr/bin/env bash

set -x
set -e
set -u

export START_PWD=$(pwd)
export WORK_DIR="$(pwd)/pulumi"

cd ${WORK_DIR}
mkdir -p ./charts

curl -L https://github.com/mastodon/chart/archive/refs/heads/main.zip -o ./charts/chart-main.zip
unzip -u -d ./charts ./charts/chart-main.zip

cd ./charts/chart-main
helm dep update

