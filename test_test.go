package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeConsuming(t *testing.T) {
	Convey("google.com", t, func() {
		days, err := checkssl("google.com")
		So(err, ShouldEqual, nil)
		So(days, ShouldBeGreaterThan, 1)
	})
	Convey("google.comm", t, func() {
		days, err := checkssl("google.comm")
		So(err, ShouldNotEqual, nil)
		So(days, ShouldBeLessThan, 0)
	})

}
