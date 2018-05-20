package provision

import (
	"fmt"
	"github.com/eosioafrica/ecte/plugin"
	"github.com/eosioafrica/ecte/environment"
	"errors"
)

var hostFileName string
var hostISO string

type Provisioner struct {

	Env     *environment.Environment
	Config 	*plugin.Configuration

	Hosts   *[]plugin.VBox

	Err error
}

/**
*
* Create new provisioner.
* Requires environment,configuration as parameters
*
**/
func New(env *environment.Environment) *Provisioner {

	return &Provisioner{

		Env: env,
		Config: &plugin.Config,
		Hosts: &[]plugin.VBox{},
		Err: nil,
	}
}


func (provisioner *Provisioner) Provision() {

	if provisioner.Err != nil { return }

	provisioner.CreateHostVMs()
}

func (provisioner *Provisioner) RunVagrant(){

	if provisioner.Err != nil { return }

	binDir := provisioner.Env.Config.Dirs.BinFull

	vagrant := plugin.Vag{

		Path: binDir,
	}

	vagrant.Provision()
}


func (provisioner *Provisioner) CreateHostVMs() {

	if provisioner.Err != nil { return }

	// Declare list of hosts to be created
	var vms []plugin.VBox

	plugin.InitialConfig(provisioner.Env)

	// Derive and fill host with configuration values
	provisioner.GenerateHostList(&vms)

	// Create all configured hosts
	for _, vm := range vms {

		vm.Destroy()

		err := vm.Create()
		if err != nil {

			provisioner.Err = vm.Err
			return
		}
	}
}

func (provisioner *Provisioner) GenerateHostList(machines *[]plugin.VBox) {

	if provisioner.Err != nil { return }

	provisionersDir := provisioner.Env.Config.Dirs.AssetsFull
	binDir := provisioner.Env.Config.Dirs.BinFull

	hostISO = fmt.Sprintf("%s/%s", provisionersDir, "ipxe.iso")

	var hosts = []plugin.VBox{}

	if len(provisioner.Config.Roles) < 1 {

		provisioner.Err = environment.WrapErrors(errors.New("No roles submitted."))
		return
	}

	for _, role := range provisioner.Config.Roles {

		if len(role.Hosts) < 1 {

			provisioner.Err = environment.WrapErrors(errors.New("No hosts submitted."))
			return
		}

		for _, host := range role.Hosts {

			hostFileName = fmt.Sprintf("%s/%s/%s.%s", binDir, host.Name, host.Name, "vdi")

			host.Interfaces = role.Interfaces
			host.Controllers = role.Controllers
			host.BootSeq = role.BootSeq

			vm := plugin.VBox{

				Name:     host.Name,
				Path:     binDir,
				OSType:   role.OsType,
				Filename: hostFileName,
				Config: plugin.MachineConfig{

					NICs: host.Networking(),
					BootSeq: host.Boot(),
					Disk: role.Disk,
					Memory: role.Memory,
					Storage: host.Storage(provisioner.Config.StorageCtls, hostISO, hostFileName),
				},
			}

			hosts = append(hosts, vm)
		}
	}
	*machines = hosts
}
