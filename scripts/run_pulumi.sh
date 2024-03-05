#!/usr/bin/env bash

set -x
set -e
set -u

# export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/pulumi-passwd
export START_PWD=$(pwd)
export WORK_DIR="pulumi"
export K_KLUSTER="mastodon.the-independent-friend.de"
export P_STACK="dev"

if [ -z ${INSTALL_AND_UPDATE_GO_MODS+x} ]
then
	INSTALL_AND_UPDATE_GO_MODS="FALSE"
	echo "var is not set '$INSTALL_AND_UPDATE_GO_MODS'"
else
	echo "var is set to '$INSTALL_AND_UPDATE_GO_MODS'"
	# INSTALL_AND_UPDATE_GO_MODS="TRUE"
fi

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

pulumi login file://$(pwd)
# pulumi up --yes
pulumi up --debug --stack ${P_STACK}

cd ${START_PWD}