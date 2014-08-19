package status

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dickeyxxx/h/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("it gets the status via HTTP", t, func() {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"status":{"Production":"red","Development":"green"},"issues":[]}`)
		}))
		url = ts.URL
		var response statusResponse
		getStatus(&response)
		So(response.Status.Production, ShouldEqual, "red")
	})

	Convey("With a status instance and context", t, func() {
		ctx := &cli.Context{}
		var stderr bytes.Buffer
		var stdout bytes.Buffer
		ctx.Stderr = &stderr
		ctx.Stdout = &stdout
		getStatus = func(response *statusResponse) {
			response.Status.Production = "gren"
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
			ctx.Args = []string{"arg1"}

			Convey("It prints the USAGE statement", func() {
				So(Run(ctx), ShouldEqual, 1)
			})
		})
	})
}
