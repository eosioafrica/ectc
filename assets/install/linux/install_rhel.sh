#!/usr/bin/env bash

# The is no default path. Valid path must be supplied
if [ -z "$1" ]
  then
    echo "Err: Missing argument. Please supply install dir as first argument."
    return 1
fi

install_bin = $1

if [ ! -d "install_bin" ]
    then
      echo "Err: Invalid path. Please supply valid dir as argument."
    return 1
fi


# -------------------------------------------------------------------
# VirtualBox installation script with Guest Additions
# -------------------------------------------------------------------

# 1. Add VirtualBox source
echo "deb http://download.virtualbox.org/virtualbox/debian `lsb_release -cs` contrib" > virtualbox.list
sudo mv virtualbox.list /etc/apt/sources.list.d/
wget -q http://download.virtualbox.org/virtualbox/debian/oracle_vbox.asc -O- | sudo apt-key add -

# 2. Install VirtualBox
sudo apt-get update
sudo apt-get install -y linux-headers-`uname -r` build-essential virtualbox-5.2 dkms


# 3. Install Guest Additions
wget http://download.virtualbox.org/virtualbox/5.2.10/VBoxGuestAdditions_5.2.10.iso

sudo mkdir /media/VBoxGuestAdditions

sudo mount -o loop,ro VBoxGuestAdditions_5.2.10.iso /media/VBoxGuestAdditions

sudo sh /media/VBoxGuestAdditions/VBoxLinuxAdditions.run

rm VBoxGuestAdditions_5.2.10.iso

sudo umount /media/VBoxGuestAdditions

sudo rmdir /media/VBoxGuestAdditions


# -------------------------------------------------------------------
# Vagrant Installation
# -------------------------------------------------------------------

temp_dir=$(mktemp -d)
curl --silent -L wget https://releases.hashicorp.com/vagrant/2.1.1/vagrant_2.1.1_x86_64.rpm -o "$temp_dir/vagrant.rpm"
pushd "$temp_dir"
sudo yum install vagrant.rpm -y
popd
rm -rf "$temp_dir"