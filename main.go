package ecte

import (

	"github.com/eosioafrica/ecte/environment"
	"fmt"
	"github.com/eosioafrica/ecte/plugin"
)

var env environment.Environment




func main() {

	dir := "/home/khosi/go/src/github.com/eosioafrica/environment/assets/provisioners"

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




