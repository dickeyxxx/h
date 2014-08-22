package util

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {
	Convey("Must", t, func() {
		Convey("Does not panic when passed nil", func() {
			So(func() {
				Must(nil)
			}, ShouldNotPanic)
		})

		Convey("Panics when passed an error", func() {
			err := errors.New("foobar")
			So(func() {
				Must(err)
			}, ShouldPanicWith, err)
		})
	})
}
