package ecte

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/spf13/viper"
)

func (env *Environment) BashDependencyInstaller() error {

	var executable string
	if _, err := os.Stat(env.SeedBashScript); err == nil {
		executable = env.SeedBashScript
	}

	if executable == "" {
		fmt.Printf("No script with name %s was found\n", env.SeedBashScript)
		return nil
	}

	arguments := []string { executable,
		viper.GetString("directories.assets.full"),
		viper.GetString("directories.bin.full"),
		viper.GetString("app.name") }

	cmd := exec.Command("bash", arguments...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
