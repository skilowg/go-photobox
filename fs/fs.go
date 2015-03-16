// Package fs manages the scanning and indexing of photos at
// a given location to display in the photobox
package fs

// List takes a path to a reachable location on the disk and returns
// the list of files and folders found there.
//
// List only returns the first level of files and folders, however, it can
// be called again with a new path if the client wishes to drill into a folder.
//
// If the given path is not reachable, List returns an error
func List(path string) ([]string, error) {

	return []string{}, nil
}
