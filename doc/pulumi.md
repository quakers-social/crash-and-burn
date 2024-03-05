PULUMI
======

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
