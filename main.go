package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/thedahv/go-photobox/lib"
)

func pathFromRequest(uri string) string {
	rx := regexp.MustCompile("\\/files\\?path=(.+)$")
	results := rx.FindStringSubmatch(uri)

	if len(results) < 2 {
		return ""
	}

	return results[1]
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	photosPath := os.Getenv("FILES")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

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
