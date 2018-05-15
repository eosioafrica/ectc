package main

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
)

/*
import "github.com/eosioafrica/ecte/ecte"

var env ecte.Environment

func main() {


}*/

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type File struct {

	Name string `mapstructure:"name"`
}

type OutputConfig struct {
	Number string		`mapstructure:"num"`
	Files []File 		`mapstructure:"files"`
}

type Role struct{

	Name   		string 			`mapstructure:"name"`
	Disk   		int    			`mapstructure:"disk"`
	Memory 		int    			`mapstructure:"memory"`
	OsType 		string 			`mapstructure:"ostype"`
	ISO    		string 			`mapstructure:"iso"`
	Hosts  		[]Host			`mapstructure:"hosts"`
	Controllers []string		`mapstructure:"controllers"`
	Interfaces  []string		`mapstructure:"interfaces"`
	Boot 		[]string		`mapstructure:"boot"`
}

type Host struct {

	Name   	string 		`mapstructure:"name"`
	Port   	[]string
	Mac 	string 		`mapstructure:"mac"`
}

type Config struct {

	Roles []Role		 `mapstructure:"roles"`
}

func main() {

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	for _, role := range c.Roles {
		fmt.Printf("name=%s\n", role.Name)
		fmt.Printf("memory=%d\n", role.Memory)
		fmt.Printf("controllers=%s\n", role.Controllers)
		fmt.Printf("interfaces=%s\n", role.Interfaces)
		fmt.Printf("boot=%s\n", role.Boot)

		for _, host := range role.Hosts {

			fmt.Printf("hostname=%s\n", host.Name)
		}
	}
}
