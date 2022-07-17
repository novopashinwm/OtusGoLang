package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, val := range env {
		if val.NeedRemove {
			os.Unsetenv(name)
			continue
		}
		os.Setenv(name, val.Value)
	}
	cmdItem := exec.Command(cmd[0], cmd[1:]...)
	cmdItem.Stdout = os.Stdout
	cmdItem.Stdin = os.Stdin
	cmdItem.Stderr = os.Stderr
	if err := cmdItem.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			returnCode = exitError.ExitCode()
		}
	}
	return returnCode

}
