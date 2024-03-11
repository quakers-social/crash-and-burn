PULUMI
======

- [PULUMI](#pulumi)
	- [MISSION](#mission)
	- [STATE](#state)
	- [INIT PULUMI PROJEKT](#init-pulumi-projekt)
	- [RUN PULUMI](#run-pulumi)
	- [GITHUB ACTION](#github-action)
	- [RUN PULUMI LOCAL BY HAND](#run-pulumi-local-by-hand)
	- [WORKING WITH SECRETS](#working-with-secrets)
		- [ADD](#add)
		- [SHOW CONFIG](#show-config)
		- [GET SECRETS](#get-secrets)
	- [TROUBLESHOOTING](#troubleshooting)
	- [TODOs](#todos)


MISSION
-------

Pulumi deeploys the Kubernetes objects.


STATE
-----

The state of Pulume is storeed as files in the git repo.


INIT PULUMI PROJEKT
-------------------

```bash
pulumi new kubernetes-go \
	--description="k3s of quakers-social crash n' burn" \
	--name="quakers-social" \
	--secrets-provider=passphrase \
	--stack="dev"
```

```bash
$ PULUMI_CONFIG_PASSPHRASE='XXXXXXXXXXX'
$ AWS_ACCESS_KEY_ID=XXXXXXXXX
$ AWS_SECRET_ACCESS_KEY=XXXXXXXXX
$ pulumi login 's3://pulumi?region=us-west-1&endpoint=us-west-1.storage.impossibleapi.net'
```

RUN PULUMI
----------

Enter the helper script:

```bash
scripts/run_pulumi.sh
```

GITHUB ACTION
-------------

Pulumi is integrated in the Github Action pipeline. All pushes check
whether the Pulumi configuration is error-free.

If a commit is ***tagged*** and this tag is pushed, then the
configuration is tested and if the test was successful, the changes
are executed by Pulumi on Kubernetes.

*Tipp:*
to find the last used tag enter this

````bash
$ git tag | sort -n | tail -n 1
````


RUN PULUMI LOCAL BY HAND
------------------------

Set the Environment variables `PULUMI_CONFIG_PASSPHRASE` or create
an passwordfile under and `${HOME}/.ssh/quakers-social/pulumi` and
set `PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/quakers-social/pulumi`
and than enter the script:

```bash
$ PULUMI_CONFIG_PASSPHRASE='XXXXXXXXXXX'
$ AWS_ACCESS_KEY_ID=XXXXXXXXX
$ AWS_SECRET_ACCESS_KEY=XXXXXXXXX
$ scripts/run_pulumi.sh
```

you will find the password you need under

`repo > settings > secrets > actions > Repository secrets`

WORKING WITH SECRETS
--------------------

### ADD

Example:

```bash
pulumi config set mongoUsername admin
pulumi config set --secret mongoPassword XXXXX
```

### SHOW CONFIG

```bash
pulumi config
```

### GET SECRETS

```bash
pulumi config --show-secrets
```


TROUBLESHOOTING
---------------

If you run the pulumi first time lokal than you have to set
```bash
$ INSTALL_AND_UPDATE_GO_MODS="true"
$ scripts/run_pulumi.sh
```

This will be downlod the needed go modules.


TODOs
-----

- Us S3 als state storage