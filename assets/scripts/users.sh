#!/usr/bin/env bash

if [ ! -f /vagrant/terraform/modules/ignition/users/user.tpl ]; then
    echo "User template file not found!"
    exit 1
fi

if [ -f /vagrant/terraform/modules/ignition/users/demo.tf ]; then
    rm /vagrant/terraform/modules/ignition/users/demo.tf
fi

ssh-keygen -q -t rsa -N "" -f /home/vagrant/.ssh/id_rsa

export PUBKEY=$(cat /home/vagrant/.ssh/id_rsa.pub)

USERDATA='$PUBKEY:$VAR2'

envsubst "$USERDATA" </vagrant/terraform/modules/ignition/users/user.tpl >/vagrant/terraform/modules/ignition/users/demo.tf



