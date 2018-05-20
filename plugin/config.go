package plugin

import (

	"github.com/spf13/viper"
	"fmt"
	"github.com/eosioafrica/ecte/environment"
)

type Role struct{

	Name   			string 			`mapstructure:"name"`
	Disk   			int    			`mapstructure:"disk"`
	Memory 			int    			`mapstructure:"memory"`
	OsType 			string 			`mapstructure:"ostype"`
	ISO    			string 			`mapstructure:"iso"`
	Hosts  			[]Host			`mapstructure:"hosts"`
	// The following can be given at both {role} and {host} levels
	Controllers 	[]string 		`mapstructure:"controllers"`
	Interfaces  	[]string 		`mapstructure:"interfaces"`
	BootSeq     	[]string 		`mapstructure:"boot"`
}

type Host struct {

	Name   			string 			`mapstructure:"name"`
	Ports   		[]string		`mapstructure:"ports"`
	Mac 			string 			`mapstructure:"mac"`
    // The following can be given at both {role} and {host} levels
	Controllers 	[]string		`mapstructure:"controllers"`
	Interfaces  	[]string		`mapstructure:"interfaces"`
	BootSeq 		[]string		`mapstructure:"boot"`

	ISO             string
	Filename		string
}

type StorageController struct {

	Name			string 			`mapstructure:"name"`
	Bus 		 	string 			`mapstructure:"bus"`
	Controller  	string 			`mapstructure:"controller"`
	Port			int 			`mapstructure:"port"`
	Device			int 			`mapstructure:"device"`
	Type			string 			`mapstructure:"type"`
	Medium			string 			`mapstructure:"medium"`
}

type Configuration struct {

	Roles 			[]Role		 			`mapstructure:"roles"`
	StorageCtls  	[]StorageController 	`mapstructure:"storage"`
}

var Config Configuration

func InitialConfig(env *environment.Environment)  {

	v := viper.New()
	v.SetConfigName("hosts")
	v.AddConfigPath(env.Config.Dirs.AssetsFull)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Named config file not found...")
		return
	}

	if err := v.Unmarshal(&Config); err != nil {

		fmt.Printf("couldn't read config: %s", err)
	}

	// TODO Reconsider location
	// Set host ipxe
	/*
	for _, role := range Config.Roles {

		for _, host := range role.Hosts {

			host.ISO = fmt.Sprintf("%s/%s", env.Config.Dirs.ProvisionersFull, "ipxe.iso")
			host.Filename = fmt.Sprintf("%s/%s/%s.%s", env.Config.Dirs.BinFull, host.Name, host.Name, "vdi")

			fmt.Println(host.ISO)
			fmt.Println(host.Filename)
		}
	}*/

}
