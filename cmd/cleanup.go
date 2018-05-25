package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

// ProvisionCmd represents client cli
var CleanupCmd = &cobra.Command{

	Use:   "cleanup",
	Short: "Removes artifacts and completly deletes the test environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		logrus.Info("Deleting test environment.")

		//		ecte := New()

		//		ecte.Run()
	},
}