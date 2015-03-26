package photobox

import (
	"fmt"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func setupFixtures() error {
	return os.Chmod("./testdata/inaccessible", 000)
}

func teardownFixtures() error {
	return os.Chmod("./testdata/inaccessible", 777)
}

func TestMain(m *testing.M) {
	// Setup
	err := setupFixtures()
	var result int

	if err != nil {
		fmt.Println(err.Error())
		result = 1
	} else {
		result = m.Run()
	}

	// Teardown
	teardownFixtures()
	os.Exit(result)
}

func TestList(t *testing.T) {
	Convey("When Listing files", t, func() {
		Convey("Given an empty path", func() {
			Convey("It should return an error", func() {
				files, err := List("")

				So(len(files), ShouldEqual, 0)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Given an accessible path", func() {
			files, err := List("./testdata/a")

			Convey("It should return no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("List should return a mix of directories and files", func() {
				So(len(files), ShouldEqual, 2)
			})

			Convey("List should exclude hidden files", func() {
				files, err := List("./testdata")
				So(err, ShouldBeNil)

				for _, file := range files {
					So(strings.HasPrefix(file.Name(), "."), ShouldBeFalse)
				}
			})
		})

		Convey("Given a non-existent path", func() {
			files, err := List("./testdata/nothing")

			Convey("It should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldResemble, "open ./testdata/nothing: no such file or directory")
			})

			Convey("Files should be empty", func() {
				So(len(files), ShouldEqual, 0)
			})
		})

		Convey("Given an inaccessible path", func() {
			files, err := List("./testdata/inaccessible")

			Convey("It should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldResemble, "open ./testdata/inaccessible: permission denied")
			})

			Convey("Files should be empty", func() {
				So(len(files), ShouldEqual, 0)
			})
		})

	})
}

func TestIsValid(t *testing.T) {
	Convey("isValid", t, func() {

		Convey("when given an empty filename", func() {

			Convey("should return false", func() {
				So(isValid(""), ShouldBeFalse)
			})

		})

		Convey("when given a filename ending in an image format", func() {

			Convey("should return true for files", func() {

				Convey("like png", func() {
					So(isValid("somefile.png"), ShouldBeTrue)
				})

				Convey("like jpg", func() {
					So(isValid("somefile.jpg"), ShouldBeTrue)
				})

				Convey("like jpeg", func() {
					So(isValid("somefile.jpeg"), ShouldBeTrue)
				})

				Convey("like tiff", func() {
					So(isValid("somefile.tiff"), ShouldBeTrue)
				})

			})

		})

		Convey("when given a filename with a valid format in the middle of the name", func() {

			Convey("like something.png.skitch", func() {

				Convey("it should return false", func() {
					So(isValid("something.png.something.skitch"), ShouldBeFalse)
				})

			})

		})

	})
}
