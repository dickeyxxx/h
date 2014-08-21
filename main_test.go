package main

import (
	"bytes"
	"testing"

	"github.com/dickeyxxx/hk/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	ctx.Stderr = &stderr
	ctx.Stdout = &stdout
	exit = func(code int) {
		panic(code)
	}
	Convey("with no arguments", t, func() {
		args = []string{"heroku"}

		Convey("it shows the help", func() {
			So(main, ShouldPanicWith, 0)
		})
	})

	Convey("with a wrong argument", t, func() {
		args = []string{"heroku", "foobar?"}

		Convey("it returns code 127", func() {
			So(main, ShouldPanicWith, 127)
		})
	})

	Convey("with a topic name", t, func() {
		args = []string{"heroku", "foo"}
		topics = []*cli.Topic{
			{Name: "foo",
				Run: func(*cli.Context) int {
					return 88
				},
			},
		}

		Convey("it runs the foo topic", func() {
			So(main, ShouldPanicWith, 88)
		})
	})
}
