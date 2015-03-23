package photobox

// Functions to manage the scanning and indexing of photos at
// a given location to display in the photobox.

import (
	"fmt"
	"io/ioutil"
)

// List takes a path to a reachable location on the disk and returns
// the list of files and folders found there.
//
// List only returns the first level of files and folders, however, it can
// be called again with a new path if the client wishes to drill into a folder.
//
// If the given path is not reachable, List returns an error.
func List(path string) ([]string, error) {
	var results []string

	if len(path) == 0 {
		return results, fmt.Errorf("path cannot be blank")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return results, err
	}

	for _, file := range files {
		results = append(results, file.Name())
	}

	return results, nil
}
