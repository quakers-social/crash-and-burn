ANSIBLE
=======

- [ANSIBLE](#ansible)
  - [MISSION](#mission)
  - [GITLAB ACTION](#gitlab-action)
  - [TROUBLESHOOTING](#troubleshooting)
  - [KNOWN ISSUE](#known-issue)
    - [Solution:](#solution)



MISSION
-------

Infrastructor as code for virtual machine(s).


GITLAB ACTION
-------------

The Ansible Playbook is primarily initiated by the *Github Action*.
See: [.github/workflows](.github/workflows)

***However, Github Action will only execute the playbook if the commit has a tag!***

However, an Ansible test run is carried out with every commit to
ensure that the playbook works at all times.

TROUBLESHOOTING
---------------

But for troubleshooting you can also run it local on your side. To do this,
you must be able to establish a working ssh connection to the VM


For a dry run with debug output enter:

```bash
ansible-playbook \
  -i ./ansible/hosts.yaml \
  ./ansible/install_and_update.yaml  \
  --diff \
  --check \
  -vvvv
```

otherwise enter:

```bash
ansible-playbook \
  -i ./ansible/hosts.yaml \
  ./ansible/install_and_update.yaml
```

You find after the playbook run the generated Kubernetes configuration in
the directory `./ansible/downloads/`


KNOWN ISSUE
-----------

The default runners of github action run in the azure cloud and have a firewall
rule that prohibits opening outgoing ssh connections (regardless of which port).

### Solution:

Using your own runner: [About self-hosted runners](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/about-self-hosted-runners)
