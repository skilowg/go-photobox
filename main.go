package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/thedahv/go-photobox/lib"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	photosPath := os.Getenv("FILES")
	canOpen, err := os.Open(photosPath)
	canOpen.Close()
	if err != nil {
		fmt.Printf("Error opening photobox: %v\n", err)
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir("public"))
	pbfs := http.FileServer(http.Dir(photosPath))

	http.Handle("/", fs)
	http.Handle("/photos/", http.StripPrefix("/photos/", pbfs))

	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		var path string
		subPath := pathFromRequest(r.RequestURI)

		if len(subPath) > 0 {
			path = photosPath + string(os.PathSeparator) + subPath
		} else {
			path = photosPath
		}

		files, err := photobox.List(path)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))

			return
		}

		type jsonFile struct {
			Name  string `json:"name"`
			IsDir bool   `json:"isDir"`
		}
		var results []jsonFile

		for _, file := range files {
			results = append(results, jsonFile{file.Name(), file.IsDir()})
		}
		jsonData, err := json.Marshal(results)
		w.Write(jsonData)
	})

	fmt.Printf("Serving files from %s on port %s\n", photosPath, port)
	http.ListenAndServe(":"+port, nil)
}

// pathFromRequest takes a URI with a URL encoded query parameter for a path
// on disk and returns that path in a format suitable for opening a file
func pathFromRequest(uri string) string {
	rx := regexp.MustCompile("\\/files\\?path=(.+)$")
	results := rx.FindStringSubmatch(uri)

	if len(results) < 2 {
		return ""
	}

	p, err := url.QueryUnescape(results[1])
	if err != nil {
		return ""
	}

	return p
}
