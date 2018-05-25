package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
	"github.com/eosioafrica/ecte/ecte"
)

var InitCmd = &cobra.Command{

	Use:   "init",
	Short: "Initialize run environment",
	Long:  `Pull seed repo from github`,
	Run: func(cmd *cobra.Command, args []string) {

		logrus.Info("Will start initializing ecte run environment.")

		ecte := ecte.New()

		ecte.Seed()

		ecte.CreateEnvironment()

		if ecte.Err != nil {

			logrus.Error("Error occured - ", ecte.Err)
		} else {

			logrus.Info("Successfully created new ecte configuration.")
			logrus.Info("Run { ecte prov } to provision new machines.")
		}
	},
}
