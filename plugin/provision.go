package plugin

import (
	"github.com/eosioafrica/ecte/ecte"
	"fmt"
	"strings"
	"strconv"
)

var hostFileName string
var hostISO string

type Provisioner struct {

	Env     *ecte.Environment
	Config 	*Configuration

	err 	error
}

/**
*
* Create new provisioner.
* Requires environment,configuration as parameters
*
**/
func New(env *ecte.Environment) *Provisioner {

	return &Provisioner{

		env,
		&Config,
		nil,
	}
}


func (provisioner *Provisioner) Provision() error {

	// Declare list of hosts to be created
	var vms []VBox

	// Derive and fill host with configuration values
	provisioner.GenerateHostList(&vms)

	// Create all configured hosts
	for _, vm := range vms {

		vm.Destroy()

		err := vm.Create()
		if err != nil {

			return vm.Err
		}
	}

	return nil
}

func (provisioner *Provisioner) RunVagrant(){

	binDir := provisioner.Env.Config.Dirs.BinFull

	vagrant := Vag{

		Path: binDir,
	}

	vagrant.Provision()
}

func (provisioner *Provisioner) GenerateHostList(machines *[]VBox) {


	assetsDir := provisioner.Env.Config.Dirs.AssetsFull
	binDir := provisioner.Env.Config.Dirs.BinFull

	var hosts = []VBox{}

	for _, role := range provisioner.Config.Roles {

		//role.ISO = fmt.Sprintf("%s/%s", assetsDir, "ipxe.iso")
		hostISO = fmt.Sprintf("%s/%s", assetsDir, "ipxe.iso")


		for _, host := range role.Hosts {

			hostFileName = fmt.Sprintf("%s/%s/%s.%s", binDir, host.Name, host.Name, "vdi")

			host.Interfaces = role.Interfaces
			host.Controllers = role.Controllers
			host.BootSeq = role.BootSeq
//			host.ISO = role.ISO

			vm := VBox{

				Name:     host.Name,
				Path:     binDir,
				OSType:   role.OsType,
				Filename: hostFileName,
//				ISO:      role.ISO,
				Config: MachineConfig{

					NICs: host.Networking(),
					BootSeq: host.Boot(),
					Disk: role.Disk,
					Memory: role.Memory,
					Storage: host.Storage(provisioner.Config.StorageCtls),
				},
			}

			hosts = append(hosts, vm)
		}

	}
	*machines = hosts
}

func (host *Host) Networking() *[]NIC {

	nics := []NIC{}

	for _ , n := range host.Interfaces {
		// TODO The interface name and mac should be more tightly coupled

		// TODO Validate n before it gets here
		stringSlice := strings.Split(n, "-")
		index, _ := strconv.Atoi(stringSlice[0])
		device := stringSlice[1]

		nic := NIC{

			Idx: index,
			Mac: host.Mac,
			Device: device,
		}

		nics = append(nics, nic)
	}

	return &nics
}

// TODO See if there is not a better way of doing this while preserving the simplicity of the .toml file
func (host *Host) Storage(controllers []StorageController) *[]Storage{

	storage := []Storage{}

	for _, ctl := range host.Controllers{
		for _, c := range controllers{

			if c.Bus == ctl {

				s := Storage{

					Bus: c.Bus,
					Device: c.Device,
					Name: c.Name,
					Controller: c.Controller,
					Type: c.Type,
					Medium: getMedium(c.Bus),
					Port: c.Port,
				}


				storage = append(storage, s)
			}
		}
	}

	return &storage
}

func getMedium (bus string) string {

	if bus == "ide"  { return hostISO }
	if bus == "sata" { return hostFileName}

	return ""
}

// TODO See if there is not a better way of doing this while preserving the simplicity of the .toml file
func (host *Host) Boot() *Boot{

	b1 := "none"
	b2 := "none"
	b3 := "none"
	b4 := "none"

	for i, b := range host.BootSeq {

		switch {
			case i == 0:
				b1 = b
			case i == 1:
				b2 = b
			case i == 2:
				b3 = b
			case i == 3:
				b4 = b
		}
	}

	return &Boot{

		Boot1: b1,
		Boot2: b2,
		Boot3: b3,
		Boot4: b4,
	}
}