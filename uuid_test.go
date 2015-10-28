package uuid

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateUUID(t *testing.T) {
	Convey("Testing UUID generation", t, func() {
		prev, err := UUID()
		So(err, ShouldBeNil)
		Convey("Checking with regular expression", func() {
			for i := 0; i < 1000; i++ {
				id, err := UUID()
				So(err, ShouldBeNil)
				So(prev, ShouldNotEqual, id)

				matched, err := regexp.MatchString(
					"[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12}", id)
				So(err, ShouldBeNil)
				So(matched, ShouldEqual, true)

			}
		})
	})

}

func TestGenerateUUIDBigData(t *testing.T) {
	Convey("Testing UUID generation (big data)", t, func() {
		prev, err := UUID()
		So(err, ShouldBeNil)
		Convey("Checking with regular expression for 4 Mio uuids", func() {
			for i := 0; i < 4000000; i++ {
				id, err := UUID()
				So(err, ShouldBeNil)
				So(prev, ShouldNotEqual, id)

				matched, err := regexp.MatchString(
					"[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12}", id)
				So(err, ShouldBeNil)
				So(matched, ShouldEqual, true)

			}
		})
	})

}
