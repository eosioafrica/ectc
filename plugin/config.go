package plugin

import (

	"github.com/spf13/viper"
	"fmt"
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

func init()  {

	v := viper.New()
	v.SetConfigName("hosts")
	v.AddConfigPath("/home/khosi/go/src/github.com/eosioafrica/ecte/assets/")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Named config file not found...")
		return
	}

	if err := v.Unmarshal(&Config); err != nil {

		fmt.Printf("couldn't read config: %s", err)
	}
}

