package main

import (
	"fmt"
	"os"

	"github.com/dickeyxxx/h/status"
)

type Command interface {
	Name() string
	Run(...string) int
}

var commands = []Command{status.New()}

var exit = os.Exit

var stderr = os.Stderr

func main() {
	if len(os.Args) < 2 {
		help()
		exit(0)
	}
	for _, command := range commands {
		if os.Args[1] == command.Name() {
			exit(command.Run(os.Args...))
		}
	}
	help()

	exit(127)
}

func help() {
	fmt.Fprintln(stderr, "USAGE: h [cmd]")
	for _, command := range commands {
		fmt.Fprintln(stderr, command.Name())
	}
}
