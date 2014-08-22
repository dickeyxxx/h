package main

import (
	"os/user"
	"path/filepath"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func homeDir() string {
	user, err := user.Current()
	must(err)
	return filepath.Join(user.HomeDir, ".hk")
}
