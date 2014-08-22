package main

import (
	"os"
	"os/exec"
)

func runRubyCli(args ...string) (int, error) {
	path := filepaths.Join(homeDir(), ".hk", "heroku.exe")
	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return 0, cmd.Run()
}
