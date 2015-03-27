package photobox

// Functions to manage the scanning and indexing of photos at
// a given location to display in the photobox.

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// imgRx will tell us if a given file matches an expected pattern
// of whitelisted formats
var imgRx *regexp.Regexp

// hiddenFile tells us whether a given file represents a hidden file
// that should be filtered from the results
const hiddenFile string = "."

func init() {
	fileEndings := []string{
		"\\.png",
		"\\.jpg",
		"\\.jpeg",
		"\\.gif",
		"\\.tiff",
	}

	imgRx = regexp.MustCompile(fmt.Sprintf("(%s)$", strings.Join(fileEndings, "|")))
}

// isValid determines whether a given filename represents a valid
// image file we want to include in our results
func isValid(fname string) bool {
	return imgRx.MatchString(fname)
}

// List takes a path to a reachable location on the disk and returns
// the list of files and folders found there.
//
// List only returns the first level of files and folders, however, it can
// be called again with a new path if the client wishes to drill into a folder.
//
// If the given path is not reachable, List returns an error.
func List(path string) ([]os.FileInfo, error) {
	var results []os.FileInfo

	if len(path) == 0 {
		return results, fmt.Errorf("path cannot be blank")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return results, err
	}

	for _, file := range files {
		n := file.Name()

		showFolder := !file.IsDir() || !strings.HasPrefix(n, hiddenFile)
		showFile := file.IsDir() || isValid(n)

		if showFolder && showFile {
			results = append(results, file)
		}
	}

	return results, nil
}
