package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	hkPath := filepath.Join(HomeDir(), ".hk", "hk")
	if runtime.GOOS == "windows" {
		hkPath = hkPath + ".exe"
	}
	exists, err := FileExists(hkPath)
	Must(err)
	if !exists {
		updateHk(hkPath)
	}
	os.Exit(run(hkPath))
}

func run(path string) int {
	cmd := exec.Command(path, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				return (status.ExitStatus())
			}
		}
	}
	return 0
}
