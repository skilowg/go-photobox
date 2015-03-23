package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPathFromRequest(t *testing.T) {
	Convey("When getting sub-path from request", t, func() {
		Convey("When the path is empty", func() {
			Convey("The result should be blank", func() {
				So(pathFromRequest(""), ShouldBeBlank)
			})
		})

		Convey("When the path has no query parameters", func() {
			Convey("The results should be blank", func() {
				So(pathFromRequest("/files"), ShouldBeBlank)
			})
		})

		Convey("When the path has the wrong query parameter", func() {
			Convey("The results should be blank", func() {
				So(pathFromRequest("/files?notarealthing=a"), ShouldBeBlank)
			})
		})

		Convey("When the path has a 'path' query parameter", func() {
			So(pathFromRequest("/files?path=subfolder"), ShouldEqual, "subfolder")
		})
	})
}
