package ecte

import (

	"github.com/eosioafrica/ecte/environment"
	"github.com/eosioafrica/ecte/seed"
	"github.com/eosioafrica/ecte/provision"
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
)


type Ecte struct {

	Seeder      	*seed.Seeder
	Environ     	*environment.Environment
	Provisioner 	*provision.Provisioner
	Err         	error
}

func New () *Ecte{

	return &Ecte{

		Seeder:      	seed.New(),
		Environ:     	environment.New(),
		Provisioner: 	nil,
		Err:         	nil,
	}
}

/*
Retrieve installation asset data from a git repo. Typically, the repo must contain amongst others,
configuration file, vagrant file for the provisioner, matchbox and terraform assets.
*/
func (ecte *Ecte) Seed (){

	if ecte.Err != nil { return }

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ecte.Seeder.ExecDir = pwd

	ecte.Err = ecte.Seeder.Seed()

	if ecte.Err ==nil { logrus.Info("Information seeding has been successful.") }
}


func (ecte *Ecte) CreateEnvironment (){

	if ecte.Err != nil { return }

	logrus.Info("Creating new ecte environment.")

	ecte.Err = ecte.Environ.Create(ecte.Seeder.AppDirCreated)

	if ecte.Err == nil { logrus.Info("New environment created at .", ecte.Seeder.AppDirCreated) }

}

func (ecte *Ecte) Provision (){

	if ecte.Err != nil { return }

	logrus.Info("Starting the provisioning of virtual machines.")

	ecte.Provisioner = provision.New(ecte.Environ)

	ecte.Provisioner.Provision()

	ecte.Err = ecte.Provisioner.Err

	if ecte.Err == nil { logrus.Info("Successfully provisioned virtual machines.") }
}

func (ecte *Ecte) Run() {

	ecte.Seed()

	ecte.CreateEnvironment()

	ecte.Provision()
}