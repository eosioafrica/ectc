package main

import (
	"github.com/eosioafrica/ecte/cmd"
	"runtime"
	"os"
	"github.com/sirupsen/logrus"
)

func main() {

	checkOS()

	cmd.RootCmd.AddCommand(cmd.ProvisionCmd)
	cmd.RootCmd.AddCommand(cmd.InitCmd)
	cmd.RootCmd.AddCommand(cmd.CleanupCmd)
	cmd.Execute()
}

func checkOS() {

	if runtime.GOOS != "linux" {

		logrus.Warning("The program ecte only supports linux.")
		logrus.Warning("Exiting... ")
		os.Exit(-1)
	}
}




