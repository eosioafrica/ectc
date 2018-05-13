package ecte

import (
	"fmt"
	"os"
	"os/exec"
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

	//cmd := exec.Command(executable, args...)
	cmd := exec.Command(executable)

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
