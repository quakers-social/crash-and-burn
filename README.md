CRASH N' BURN
=============

![example workflow](https://github.com/quakers-social/crash-and-burn/actions/workflows/setup_vm.yaml/badge.svg)
![example workflow](https://github.com/quakers-social/crash-and-burn/actions/workflows/setup_kubernetes.yaml/badge.svg)
![example workflow](https://github.com/quakers-social/crash-and-burn/actions/workflows/image_build.yaml/badge.svg)


MISSION
-------

This is the infrastructure as code for a test environment of Mastodon.


TECH STACK
----------

- ***Ansible:*** for set up of the virtual machine and Kubernetes
- ***Pulumi:*** for manage the Kubernetes Workload (Resources)
- ***Kubernetes Operators:*** for managing let's encrypt, PostgreSQL and Mastodon
- ***S3:*** for storage Pulumi state, Mastodon media files and Postgres backups
- ***GitHub Actions:*** triggering Ansible and Pulumi

DOKUMENTATION
-------------

See: [doc/_INDEX.md](doc/_INDEX.md)
