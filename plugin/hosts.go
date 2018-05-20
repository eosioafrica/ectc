package plugin

import (

	"strings"
	"strconv"
)



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
func (host *Host) Storage(controllers []StorageController, iso string, file string) *[]Storage{

	storage := []Storage{}

	for _, ctl := range host.Controllers{
		for _, c := range controllers{

			if c.Bus == ctl {

				var medium string

				if c.Bus == "ide"  { medium = iso }
				if c.Bus == "sata" { medium = file }

				s := Storage{

					Bus: c.Bus,
					Device: c.Device,
					Name: c.Name,
					Controller: c.Controller,
					Type: c.Type,
					Medium: medium,
					Port: c.Port,
				}

				storage = append(storage, s)
			}
		}
	}

	return &storage
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