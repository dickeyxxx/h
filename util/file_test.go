package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFile(t *testing.T) {
	Convey("FileExists", t, func() {
		Convey("returns true if file exists", func() {
			exists, err := FileExists("./file_test.go")
			So(exists, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("returns false if file does not exist", func() {
			exists, err := FileExists("./foobar.go")
			So(exists, ShouldBeFalse)
			So(err, ShouldBeNil)
		})
	})
}
