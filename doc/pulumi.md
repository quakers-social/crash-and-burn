PULUMI
======


- [PULUMI](#pulumi)
	- [MISSION](#mission)
	- [STATE](#state)
	- [INIT PULUMI PROJEKT](#init-pulumi-projekt)
	- [RUN PULUMI](#run-pulumi)
	- [GITHUB ACTION](#github-action)
	- [RUN PULUMI LOCAL BY HAND](#run-pulumi-local-by-hand)
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
$ pulumi login file://$(pwd)
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

RUN PULUMI LOCAL BY HAND
------------------------

Set the Environment variables `PULUMI_CONFIG_PASSPHRASE` and
enter the halber script:

```bash
$ PULUMI_CONFIG_PASSPHRASE='XXXXXXXXXXX'
$ scripts/run_pulumi.sh
```

you will find the password you need under

`repo > settings > secrets > actions > Repository secrets`

TODOs
-----

- Us S3 als state storage