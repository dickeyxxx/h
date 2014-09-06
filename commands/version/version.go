package version

import (
	"runtime"

	"github.com/dickeyxxx/hk/cli"
)

var Topic = &cli.Topic{
	Name: "version",
	Run:  Run,
}

func Run(ctx *cli.Context) int {
	ctx.Printf("hk/%s (%s-%s) %s\n", ctx.Version, runtime.GOARCH, runtime.GOOS, runtime.Version())
	return 0
}
