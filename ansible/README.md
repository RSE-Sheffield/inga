# Deploying INGA using Ansible

## Overview

[Ansible][ansible] configuration to deploy an INGA service to:
 - A remote server (e.g. a University of Sheffield-provided VMWare VM)
 - An ephemeral local [VirtualBox][virtualbox] VM (created using [Vagrant][vagrant]; for testing purposes only)

## Usage

### Local testing using Vagrant and VirtualBox

 1. Ensure the VirtualBox host machine has sufficient RAM to run the guest VM
    and that that machines and the machines it typically talks to 
    don't have addresses on the 192.168.34.0/24 network.
 1. Install some pre-requisites:
      * VirtualBox
      * Vagrant
      * Ansible
 1. Clone this repository and `cd` into the repository directory
 1. Build the INGA binary
 1. `cd` into the `ansible` directory
 1. Run `vagrant up` to:
      - Download a particular Ubuntu 18.04 Vagrant image (if not already cached locally)
      - Create and start a VirtulBox VM using that image
      - Provision that VM using particular Ansible configuration (a *playbook*, a *role*, an *inventory*, global *variables*)
 5. INGA should then be available at `http://192.168.34.10`

You can also SSH into this VM using:

```sh
vagrant ssh
```

To destroy this VM when you're finished using it:

```sh
vagrant destroy
```

If you make any changes to the Ansible/Vagrant config you then need to:

```sh
vagrant up  # if you haven't done so already
vagrant provision
```

## Deploying to persistent remote VMs

 1. Create/request a new VM e.g. on the IT dept's VMWare estate.
 2. Ensure Ansible is installed on your local machine.
 1. Clone this repository and `cd` into the repository directory
 1. Build the INGA binary
 1. `cd` into the `ansible` directory
 1. Provision the VM using:

    ```sh
    ansible-playbook  --ask-become-pass --become --inventory=inventory.ini --limit=dev playbook.yml
    ```

    or, to check only (will not provision if using `--check`)

    ```sh
    ansible-playbook --check --diff --ask-become-pass --become --inventory=inventory.ini playbook.yml
    ```

## Notes on structure of this Ansible config

* `inventory.ini`: defines and groups the hosts that we want to configure.
  We currently only have one hostgroup containing one host 
  (deploying to local VMS via Vagrant is a special case as
  it manages its inventory in a different way);
* `playbook.yml`: specifies which roles should be applied to which hostgroups.
* `Vagrantfile`: config needed by Vagrant to create a VM using a specific image, configure networking and provision the VM using Ansible.
* `roles/inga/`: the one Ansible role we want to apply to our VM(s)
* `roles/inga/tasks/main.yml`: The set of tasks to be undertaken if this role is applied.  
  These are processed in order.
  Some trigger `handlers` (via per-task `notify` definitions) when there is a state change.
* `roles/inga/handlers/main.yml`: Handlers are tasks that triggered by a state change.
  For a given *play* (playbook run) each handler is executed at most once after all non-handler tasks have been executed.
* `roles/inga/defaults/main.yml`: tasks, handlers and templates (see below) are parameterised by variables.
  Some are defined here to provide defaults for the role.  Variables can also be defined in playbooks, inventories, Vagrantfiles, on the command-line and in several other places.
  These definitions may override role defaults.
* `roles/inga/templates/`: Jinja2 template(s) that are used to create file(s) on the target machines.

## Tips for developers

### Ansible and Vagrant syntax checking

Enable basic validation by creating a file called `.git/hooks/pre-commit` within your local copy of this repository.  
This file should contain:

```sh
#!/bin/bash
set -eu
ansible-playbook --syntax-check -i inventory.ini playbook.yml && vagrant validate
```

After creating this you need to run:

```sh
chmod +x .git/hooks/pre-commit` 
```

The syntax of the Ansible and Vagrant config in this repo will then be checked the next time you commit to the repository!


[ansible]: https://www.ansible.com/
[vagrant]: https://www.vagrantup.com/
[virtualbox]: https://www.virtualbox.org/
