PULUMI
======


- [PULUMI](#pulumi)
	- [MISSION](#mission)
	- [STATE](#state)
	- [INIT PULUMI PROJEKT](#init-pulumi-projekt)
	- [RUN PULUMI](#run-pulumi)
	- [GITHUB ACTION](#github-action)
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

Next...

TODOs
-----

- Us S3 als state storage