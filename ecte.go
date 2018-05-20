package ecte

import (
	"github.com/eosioafrica/ecte/environment"
	"github.com/eosioafrica/ecte/seed"
	"github.com/eosioafrica/ecte/provision"
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


func (ecte *Ecte) Seed (){

	if ecte.Err != nil { return }

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ecte.Seeder.ExecDir = pwd

	ecte.Err = ecte.Seeder.Seed()
}


func (ecte *Ecte) CreateEnvironment (){

	if ecte.Err != nil { return }

	ecte.Err = ecte.Environ.Create(ecte.Seeder.AppDirCreated)
}

func (ecte *Ecte) Provision (){

	if ecte.Err != nil { return }

	ecte.Provisioner = provision.New(ecte.Environ)

	ecte.Provisioner.Provision()

	ecte.Err = ecte.Provisioner.Err
}

func (ecte *Ecte) Run() {

	ecte.Seed()

	ecte.CreateEnvironment()

	ecte.Provision()
}