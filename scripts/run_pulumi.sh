#!/usr/bin/env bash

set -x
set -e
set -u

export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/quakers-social/pulumi
export START_PWD=$(pwd)
export WORK_DIR="$(pwd)/pulumi"
export K_KLUSTER="mastodon.the-independent-friend.de"
export P_STACK="dev"

# Don't forgot to set:
# AWS_ACCESS_KEY_ID=XXXXXXXXX
# AWS_SECRET_ACCESS_KEY=XXXXXXXXX

if [ -z ${INSTALL_AND_UPDATE_GO_MODS+x} ]
then
	INSTALL_AND_UPDATE_GO_MODS="FALSE"
	echo "var is not set '$INSTALL_AND_UPDATE_GO_MODS'"
else
	echo "var is set to '$INSTALL_AND_UPDATE_GO_MODS'"
	# INSTALL_AND_UPDATE_GO_MODS="TRUE"
fi

cd ${WORK_DIR}
mkdir -p ./charts

curl -L https://github.com/mastodon/chart/archive/refs/heads/main.zip -o ./charts/chart-main.zip
unzip -u -d ./charts ./charts/chart-main.zip

cd ./charts/chart-main
helm dep update
cd ${WORK_DIR}

kubectx ${K_KLUSTER}
kubectl get node

go mod tidy
if [ ${INSTALL_AND_UPDATE_GO_MODS} == "TRUE" ]
then
	go get -u -t -x
else
	echo "don't update go modules."
fi

pulumi login 's3://pulumi?region=us-west-1&endpoint=us-west-1.storage.impossibleapi.net'
# pulumi up --yes
pulumi up --debug --stack ${P_STACK}

cd ${START_PWD}