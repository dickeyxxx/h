package main

import "os/user"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func homeDir() string {
	user, err := user.Current()
	must(err)
	return user.HomeDir
}
