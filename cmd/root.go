package cmd

import (
"fmt"
"os"

"github.com/spf13/cobra"
"github.com/spf13/viper"
"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

var seedRepo string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ecte",
	Short: "Eos.io clustered test environment.",
	Long: `Ecte is used to create repeatable test and development environments for the EOS.IO project.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		logrus.Info("Starting ecte.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&seedRepo, "seed", "", "github seed repo (default is https://github.com/eosioafrica/seed.git)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if seedRepo != "" {
		// Use config file from the flag.
		viper.SetConfigFile(seedRepo)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ecte")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}