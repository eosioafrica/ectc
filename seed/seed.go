package seed

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"os"
	"gopkg.in/src-d/go-git.v4"
	"github.com/eosioafrica/ecte/utils"
)

type Seed struct {

	Git 	string		`mapstructure:"git"`
}

type SeedConfiguration struct {

	Seed			Seed 		`mapstructure:"seed"`
	Destination 	string
}

var SeedConfig SeedConfiguration

type Seeder struct {

	ExecDir			string

	Config SeedConfiguration

	AppDirCreated 	string
	// The resulting new assets dir that is to be passed to enviroment package
	AssetsDirCreated string
	Err error
}

func New() *Seeder {

	return &Seeder{

		Config: SeedConfig,
		Err: nil,
	}
}


func (seeder *Seeder) StartSeed (){

	v := viper.New()
	v.SetConfigName("EcteSeed")
	v.AddConfigPath(seeder.ExecDir)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("EcteSeed not found.")
		panic(err)
	}

	if err := v.Unmarshal(&SeedConfig); err != nil {

		fmt.Printf("couldn't read config: %s", err)
		panic(err)
	}
}

// Returns an error object  and the create assets dir.

func (seeder *Seeder) Seed () error  {

	seeder.StartSeed()

	//put a starting logger here

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ecteDir := fmt.Sprintf("%s/%s", home, ".ecte")
	utils.RemoveDirIfExist(ecteDir)
	assetsDir := fmt.Sprintf("%s/%s", home, ".ecte/assets")
	fmt.Println("Setting up environment at : ", ecteDir)

	utils.CreateDirIfNotExist(assetsDir)

	seeder.Config.Destination = assetsDir

	seeder.DoGitPull()

	seeder.AppDirCreated = ecteDir
	seeder.AssetsDirCreated = seeder.Config.Destination

	return seeder.Err
}

func (seeder *Seeder) DoGitPull(){

	if seeder.Err != nil { return }

	_, err := git.PlainClone(seeder.Config.Destination, false, &git.CloneOptions{
		URL:      SeedConfig.Seed.Git,
		Progress: os.Stdout,
	})

	if err !=nil {

		seeder.Err = err
	}

	return
}

