package main

import (
	"log"
	"os"
	"os/user"
	"syscall"
	"time"

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
		runRubyCli()
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

func runRubyCli() {
	env := os.Environ()
	err := syscall.Exec(homeDir()+"/.heroku/client/bin/heroku", ctx.Args, env)
	must(err)
}

func homeDir() string {
	user, err := user.Current()
	must(err)
	return user.HomeDir
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func hkDir() string {
	return homeDir() + "/.hk"
}

func shouldAutoupdate() bool {
	if f, err := os.Stat(hkDir() + "/autoupdate"); err == nil {
		return f.ModTime().Add(4 * time.Second).Before(time.Now())
	}
	return true
}

func autoupdate() {
	err := os.MkdirAll(hkDir(), 0777)
	must(err)
	file, err := os.OpenFile(hkDir()+"/autoupdate", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	must(err)
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("checking for update")
}
