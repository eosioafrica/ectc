package ecte

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

/*
1. Ensure folder exists
2. Get asset from github and store in asset folder
3. Do chmod
4. Start Execution of script
*/

type Seeder struct {

	env *Environment
	err error
}

func (env *Environment) SeederHandler ()  {

	seeder := Seeder{

		env: env,
		err: nil,
	}

	//put a starting logger here

	seeder.CheckIfDownloadDirExists()
	seeder.RunBashDependencyDownload()
	seeder.RunBashDependencyInstallation()

	//put a finished logger here
}

func (seed *Seeder) CheckIfDownloadDirExists() {

	if _, err := os.Stat(viper.GetString("directories.assets.full")); os.IsNotExist(err) {

		seed.err = WrapErrors(ErrAssetDirDoesNotExist)
	}
}

func (seed *Seeder) RunBashDependencyDownload() {

	if seed.err != nil { return }

	if err := seed.env.DownloadSeedBashInstallAsset(); err != nil{

		seed.err = WrapErrors(ErrDownloadingBashSeedAsset, err)
	}
}

func (seed *Seeder) RunBashDependencyInstallation() {

	if seed.err != nil { return }

	if err := seed.env.BashDependencyInstaller(); err != nil{

		seed.err = WrapErrors(ErrRunBashDependencyInstallation, err)
	}

	fmt.Println(InfoRunBashDependencyInstallationSuccess)
}