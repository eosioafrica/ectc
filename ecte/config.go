package ecte

import (

	"github.com/spf13/viper"
	"os"
	"fmt"
	"github.com/mitchellh/go-homedir"
)

type Application struct{

	Name   			string 			`mapstructure:"name"`
	Seed   			string    		`mapstructure:"bashseed"`
}

type Bash struct{

	MkDir   		string 			`mapstructure:"mkdir"`
	UserAdd   		string    		`mapstructure:"useradd"`
	UserMod   		string 			`mapstructure:"usermod"`
	UserDel   		string    		`mapstructure:"userdel"`
	GroupAdd   		string 			`mapstructure:"groupadd"`
	GroupDel   		string    		`mapstructure:"groupdel"`
	Echo   			string 			`mapstructure:"echo"`
	RM   			string    		`mapstructure:"rm"`
}

type Directories struct{

	Content   			string 			`mapstructure:"content"`
	VirtualBox   		string    		`mapstructure:"virtualbox"`
	Assets   			string 			`mapstructure:"assets"`
	Bin		   			string    		`mapstructure:"bin"`

	VirtualBoxFull 		string
	AssetsFull 			string
	BinFull	   			string
	ProvisionersFull	string
}


type EnviroConfig struct {

	App 			Application		 		`mapstructure:"application"`
	Bash  			Bash 					`mapstructure:"bash"`
	Dirs  			Directories 			`mapstructure:"directories"`

	PWD				string
	Home  			string
}

var EnvConfig EnviroConfig

func init(){

	v := viper.New()
	v.SetConfigName("environment")
	v.AddConfigPath("/home/khosi/go/src/github.com/eosioafrica/ecte/assets/")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Named config file not found...")
		return
	}

	if err := v.Unmarshal(&EnvConfig); err != nil {

		fmt.Printf("couldn't read config: %s", err)
	}
	SetServiceDefault()
}


func SetServiceDefault()  {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	EnvConfig.Home = home

	fmt.Println(InfoEnvironmentBaseDirectory, EnvConfig.Home)

	EnvConfig.Dirs.VirtualBoxFull = fmt.Sprintf("%s/%s", EnvConfig.Home, EnvConfig.Dirs.VirtualBox)
	EnvConfig.Dirs.AssetsFull = fmt.Sprintf("%s/%s", EnvConfig.Home, EnvConfig.Dirs.Assets)
	EnvConfig.Dirs.BinFull = fmt.Sprintf("%s/%s", EnvConfig.Home, EnvConfig.Dirs.Bin)
	EnvConfig.Dirs.ProvisionersFull = fmt.Sprintf("%s/%s", EnvConfig.Dirs.AssetsFull, "provisioners")

}


func(env *Environment) GetDirsToCreate() []string {

	dirs := []string{EnvConfig.Dirs.VirtualBoxFull,
		EnvConfig.Dirs.AssetsFull, EnvConfig.Dirs.BinFull}

	return dirs
}



