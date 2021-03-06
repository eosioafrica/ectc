package environment

import (
	"fmt"
	"os"
	"os/exec"
)

func (env *Environment) BashDependencyInstaller() {

	if env.Err != nil { return }

	var executable string
	if _, err := os.Stat(env.SeedBashScript); err == nil {
		executable = env.SeedBashScript
	}

	if executable == "" {
		fmt.Printf("No script with name %s was found\n", env.SeedBashScript)
		env.Err = WrapErrors(ErrInstallBashScriptNotFound)
	}

	arguments := []string { executable,
		env.Config.Dirs.AssetsFull,
		env.Config.Dirs.BinFull,
		env.Config.App.Name }

	cmd := exec.Command("bash", arguments...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	err := cmd.Run()
	if err != nil {
		env.Err = WrapErrors(ErrRunBashDependencyInstallation, err)
	}

	return
}
