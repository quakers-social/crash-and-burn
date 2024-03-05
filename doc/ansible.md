ANSIBLE
=======

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