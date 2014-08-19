package status

import (
	"bytes"
	"testing"

	"github.com/dickeyxxx/h/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("With a status instance and context", t, func() {
		ctx := cli.NewContext([]string{})
		var stderr bytes.Buffer
		var stdout bytes.Buffer
		ctx.Stderr = &stderr
		ctx.Stdout = &stdout
		getStatus = func(response *statusResponse) {
			response.Status.Production = "green"
		}

		Convey("it gets a foobarred production status", func() {
			getStatus = func(response *statusResponse) {
				response.Status.Production = "foobarred."
			}
			Run(ctx)
			So(stdout.String(), ShouldContainSubstring, "Production:   foobarred.")
		})

		Convey("it gets a green production status", func() {
			Run(ctx)
			So(stdout.String(), ShouldContainSubstring, "Production:   No known issues at this time.")
		})

		Convey("With one argument", func() {
			ctx := cli.NewContext([]string{"arg1"})

			Convey("It prints the USAGE statement", func() {
				So(Run(ctx), ShouldEqual, 1)
			})
		})
	})
}
