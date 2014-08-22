package main

import (
	"os"
	"os/exec"
	"syscall"
)

func runRubyCli(args ...string) int {
	args = append([]string{homeDir() + "/.heroku/client/bin/heroku"}, args...)
	cmd := exec.Command("ruby", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				return (status.ExitStatus())
			}
		}
		panic(err)
	}
	return 0
}
