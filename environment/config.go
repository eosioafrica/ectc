package environment

import (

	"os"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Application struct{

	Name   			string 			`mapstructure:"name"`
	Seed   			string    		`mapstructure:"bashseed"`
	User  			string    		`mapstructure:"user"`
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

	// Full formed paths that map/house to relative folder paths
	VirtualBoxFull   string
	AssetsFull       string
	BinFull          string
	ProvisionersFull string
	AppPathFull      string
}


type EnviroConfig struct {

	SourcePath		string

	App 			Application		 	`mapstructure:"application"`
	Bash  			Bash 				`mapstructure:"bash"`
	Dirs  			Directories 		//`mapstructure:"directories"`

	PWD				string
	Home  			string
}

var EnvConfig EnviroConfig

func (env *Environment) InitConfig()  {

	assetsPath := fmt.Sprintf("%s/%s", EnvConfig.SourcePath, "assets")

	v := viper.New()
	v.SetConfigName("environment")
	v.AddConfigPath(assetsPath)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Named config file not found...")
		panic(err)
	}

	if err := v.Unmarshal(&EnvConfig); err != nil {

		fmt.Printf("couldn't read config: %s", err)
	}

	EnvConfig.Dirs.AppPathFull = env.Config.SourcePath
	EnvConfig.Dirs.AssetsFull = assetsPath
	SetServiceDefault()
}

func SetServiceDefault()  {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	EnvConfig.Home = home

	fmt.Println(InfoEnvironmentBaseDirectory, EnvConfig.SourcePath)

	EnvConfig.Dirs.ProvisionersFull = fmt.Sprintf("%s/%s", EnvConfig.Dirs.AssetsFull, "provisioners")
	EnvConfig.Dirs.VirtualBoxFull = fmt.Sprintf("%s/%s", EnvConfig.Dirs.AppPathFull, "virtualbox")
	EnvConfig.Dirs.BinFull = fmt.Sprintf("%s/%s", EnvConfig.Dirs.AppPathFull, "bin")
}

func(env *Environment) GetDirsToCreate() []string {

	// .ecte/assets/provisioners already exists. Only two more branches need to be created
	dirs := []string{EnvConfig.Dirs.VirtualBoxFull, EnvConfig.Dirs.BinFull}

	return dirs
}



