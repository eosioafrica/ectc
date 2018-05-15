package ecte

import (

	"github.com/spf13/viper"
	"os"
	"fmt"
)


func init(){

	viper.SetDefault("app.name", "ECTE")

	viper.SetDefault("CMDMkDir", "/bin/mkdir")
	viper.SetDefault("CMDUserAdd", "/usr/sbin/useradd")
	viper.SetDefault("CMDUserMod", "/usr/sbin/usermod")
	viper.SetDefault("CMDUserDel", "/usr/sbin/userdel")
	viper.SetDefault("CMDGroupAdd", "/usr/sbin/groupadd")
	viper.SetDefault("CMDGroupDel", "/usr/sbin/groupdel")

	viper.SetDefault("CMDEcho", "/bin/echo")
	viper.SetDefault("CMDRM", "/bin/rm")

	viper.SetDefault("directories.content", ".ecte")
	viper.SetDefault("directories.virtualbox", ".ecte/virtualbox")
	viper.SetDefault("directories.assets", ".ecte/assets")
	viper.SetDefault("directories.bin", ".ecte/bin")

	viper.SetDefault("assets.bash_seed",
		"https://raw.githubusercontent.com/khosimorafo/assets/master/get_dependency_installer.sh")

	SetServiceDefault()
}

func (env *Environment) GetDefaultConfig() string {

	return viper.GetString("CMDRM")
}

func SetServiceDefault()  {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(InfoEnvironmentBaseDirectory, pwd)

	viper.SetDefault("directories.pwd", pwd)

	vbox := fmt.Sprintf("%s/%s", viper.GetString("directories.pwd"), viper.GetString("directories.virtualbox"))
	assets := fmt.Sprintf("%s/%s", viper.GetString("directories.pwd"), viper.GetString("directories.assets"))
	bin := fmt.Sprintf("%s/%s", viper.GetString("directories.pwd"), viper.GetString("directories.bin"))

	// Set fully qualified folder path for app ready use.
	viper.SetDefault("directories.virtualbox.full", vbox)
	viper.SetDefault("directories.assets.full", assets)
	viper.SetDefault("directories.bin.full", bin)
}




func(env *Environment) GetDirsToCreate() []string {

	elements := []string{ viper.GetString("directories.virtualbox.full"),
		viper.GetString("directories.assets.full"), viper.GetString("directories.bin.full") }

	return elements
}

func(env *Environment) GetAssetsFullPath() ( string, error ) {

	return "/home/khosi/go/src/github.com/eosioafrica/ecte/assets", nil
	//return viper.GetString("directories.assets.full"), nil
}

func(env *Environment) GetVirtualBoxFullPath() ( string, error ) {

	return viper.GetString("directories.virtualbox.full"), nil
}

func(env *Environment) GetBinFullPath() ( string, error ) {

	return viper.GetString("directories.bin.full"), nil
}





