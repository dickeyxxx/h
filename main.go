package main

import (
	"os"

	"github.com/dickeyxxx/hk/cli"
	"github.com/dickeyxxx/hk/commands/status"
	"github.com/dickeyxxx/hk/commands/version"
)

var topics = []*cli.Topic{
	status.Topic,
	version.Topic,
}

var ctx = &cli.Context{
	Stdout:  os.Stdout,
	Stderr:  os.Stderr,
	Version: VERSION,
}
var exit = os.Exit
var args = os.Args
var VERSION string

func main() {
	ctx.Args = args
	code := 127
	if len(args) > 1 {
		code = runCommand(ctx, topics)
	}
	//if code == 127 {
	//var err error
	//code, err = runRubyCli(os.Args[1:]...)
	//must(err)
	//}
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
