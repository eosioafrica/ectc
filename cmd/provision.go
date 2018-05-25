package cmd

import (

	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

// ProvisionCmd represents client cli
var ProvisionCmd = &cobra.Command{

	Use:   "prov",
	Short: "Provision new test environment",
	Long:  `Create and provisioner and nodes as per matchbox assets`,
	Run: func(cmd *cobra.Command, args []string) {

		logrus.Info("Provisioning new environment.")

		//		ecte := New()

		//		ecte.Run()
	},
}