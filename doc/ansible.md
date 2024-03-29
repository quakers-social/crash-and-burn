ANSIBLE
=======

- [ANSIBLE](#ansible)
  - [MISSION](#mission)
  - [GITLAB ACTION](#gitlab-action)
  - [TROUBLESHOOTING](#troubleshooting)
  - [UNINSTALL K3S](#uninstall-k3s)
  - [ENCRYPTING CONTENT WITH ANSIBLE VAULT](#encrypting-content-with-ansible-vault)
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
  --vault-password-file  ./ansible/.ansible-password-file \
  -i ./ansible/hosts.yaml \
  ./ansible/install_and_update.yaml  \
  --diff \
  --check \
  -vvvv
```

otherwise enter:

```bash
ansible-playbook \
  --vault-password-file  ./ansible/.ansible-password-file \
  -i ./ansible/hosts.yaml \
  ./ansible/install_and_update.yaml
```

You find after the playbook run the generated Kubernetes configuration in
the directory `./ansible/downloads/`


UNINSTALL K3S
-------------

For uninstall the k3s cluster enter:

```bash
ansible-playbook \
  -i ./ansible/hosts.yaml \
  ./ansible/run_uninstall.yaml  \
  --diff \
  --check
```

ENCRYPTING CONTENT WITH ANSIBLE VAULT
-------------------------------------

The origin documentation [see here](https://docs.ansible.com/ansible/latest/vault_guide/vault_encrypting_content.html)


Encrypted variables:

```bash
ansible-vault encrypt_string --vault-password-file  ./ansible/.ansible-password-file  'foobar' --name 'the_secret'
```



KNOWN ISSUE
-----------

The default runners of github action run in the azure cloud and have a firewall
rule that prohibits opening outgoing ssh connections (regardless of which port).

### Solution:

Using your own runner: [About self-hosted runners](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/about-self-hosted-runners)

See issue: [[CI/CD] We use our own Github Action Runners #20](https://github.com/quakers-social/production-tracker/issues/20)
