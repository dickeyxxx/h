package main

import "os/user"

func HomeDir() string {
	user, err := user.Current()
	Must(err)
	return user.HomeDir
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
