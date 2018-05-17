package main

import (

	"github.com/eosioafrica/ecte/ecte"
	"fmt"
	"github.com/eosioafrica/ecte/plugin"
)

var env ecte.Environment

func main() {

	dir := "/home/khosi/go/src/github.com/eosioafrica/ecte/assets/provisioners"


	vagrant := plugin.Vag{

		Path: dir,
	}

	vagrant.Provision()

	if vagrant.Err != nil {

		fmt.Println(vagrant.Err)
		return
	}


	fmt.Println("Cluster provisioning has been successful.............")

}


