package plugin


import (
	"log"

	"github.com/koding/vagrantutil"
	"fmt"
)


// Vagrant represents a VirtualBox machine
type Vag struct {

	Name   string
	Path   string 		// Install path

	Err 		error
}

func (vag *Vag) Provision () {

	vagrant, _ := vagrantutil.NewVagrant(vag.Path)

	vagrant.Create(`# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile LEAVE AS IS!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  hostname = "provisioner-1.infra.local"
  config.vm.box = "ubuntu/xenial64"
  config.vm.hostname = hostname
  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
    v.cpus = 1
  end

  internal_ip = "192.168.1.1"
  config.vm.network "private_network", ip: internal_ip,
    virtualbox__intnet: true

  config.vm.provision "shell", path: "scripts/provision.sh",
    args: [ internal_ip ]

end
`)

	status, _ := vagrant.Status() // prints "NotCreated"

	if status.String() != "NotCreated" {

		fmt.Println("A provisioned environment exists. Please remove it should you want to continue. ")
		return
	}

	fmt.Println("Creating provisioner box. This will take a while... ")

	// starts the box
	output, _ := vagrant.Up()

	// print the output
	for line := range output {
		log.Println(line)
	}

	fmt.Println("Box provisioning has been completed. ")
	// starts the box
	output1, err := vagrant.SSH("ls -al")

	// print the output
	for line := range output1 {
		log.Println("SSH : ", line)
	}

	if err != nil {

		log.Println("SSH error : ", err)
	}

	// stop/halt the box
	//vagrant.Halt()

	// destroy the box
	//vagrant.Destroy()
}