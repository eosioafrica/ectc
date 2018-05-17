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

func NewVag (path string) *Vag{

	return &Vag{

		Path: path,
	}
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
	fmt.Println(status)

	// starts the box
	output, _ := vagrant.Up()

	// print the output
	for line := range output {
		log.Println(line)
	}


	// starts the box
	//output1, _ := vagrant.SSH("ls -al /home")

	// print the output
	//for line := range output1 {
	//	log.Println(line)
	//}

	// stop/halt the box
	//vagrant.Halt()

	// destroy the box
	//vagrant.Destroy()
}