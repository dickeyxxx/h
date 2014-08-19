package main

import (
	"os"

	"github.com/dickeyxxx/h/cli"
	"github.com/dickeyxxx/h/status"
)

var topics = []*cli.Topic{
	status.Topic,
}

var ctx = cli.NewContext(os.Args[2:])
var exit = os.Exit
var args = os.Args

func main() {
	if len(args) < 2 {
		help()
		exit(0)
	}
	ctx.Args = args[2:]
	for _, topic := range topics {
		if args[1] == topic.Name {
			exit(topic.Run(ctx))
		}
	}
	help()

	exit(127)
}

func help() {
	ctx.ErrPrintln("USAGE: h [cmd]")
	for _, topic := range topics {
		ctx.ErrPrintln(topic.Name)
	}
}
