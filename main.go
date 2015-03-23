package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

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

		paths, err := photobox.List(path)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(strings.Join(paths, ",")))
		}

	})

	fmt.Printf("Serving files from %s on port %s\n", photosPath, port)
	http.ListenAndServe(":"+port, nil)
}
