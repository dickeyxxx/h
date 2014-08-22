package main

import (
	"os"

	"github.com/dickeyxxx/hk/cli"
	"github.com/dickeyxxx/hk/status"
)

var topics = []*cli.Topic{
	status.Topic,
}

var ctx = &cli.Context{Stdout: os.Stdout, Stderr: os.Stderr}
var exit = os.Exit
var args = os.Args

func main() {
	if shouldAutoupdate() {
		autoupdate()
	}
	ctx.Args = args
	code := 127
	if len(args) > 1 {
		code = runCommand(ctx, topics)
	}
	if code == 127 {
		var err error
		code, err = runRubyCli(os.Args[1:]...)
		if err != nil {
			panic(err)
		}
	}
	exit(code)
}

func runCommand(ctx *cli.Context, topics []*cli.Topic) int {
	for _, topic := range topics {
		if ctx.Args[1] == topic.Name {
			return topic.Run(ctx)
		}
	}
	return 127
}
