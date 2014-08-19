package cli

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	Convey("with a context", t, func() {
		ctx := &Context{}
		var stderr, stdout bytes.Buffer
		ctx.Stderr = &stderr
		ctx.Stdout = &stdout

		Convey(".Print", func() {
			ctx.Print("foo")
			So(stdout.String(), ShouldEqual, "foo")
		})

		Convey(".Println", func() {
			ctx.Println("foo")
			So(stdout.String(), ShouldEqual, "foo\n")
		})

		Convey(".ErrPrint", func() {
			ctx.ErrPrint("foo")
			So(stderr.String(), ShouldEqual, "foo")
		})

		Convey(".ErrPrintln", func() {
			ctx.ErrPrintln("foo")
			So(stderr.String(), ShouldEqual, "foo\n")
		})
	})
}
