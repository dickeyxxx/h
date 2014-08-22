package main

import (
	"os"
	"os/exec"
	"syscall"
)

func runRubyCli(args ...string) int {
	cmd := exec.Command(homeDir()+"\\.hk\\heroku.exe", args...)
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
