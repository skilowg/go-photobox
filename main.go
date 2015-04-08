package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/jaschaephraim/lrserver"
	"github.com/thedahv/go-photobox/lib"
	"gopkg.in/fsnotify.v1"
)

var appMode string

func init() {
	appMode = os.Getenv("APPMODE")
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	// Serve homepage
	t, err := template.ParseFiles("./webapp/views/index.html")

	if err != nil {
		var errorMsg string

		w.WriteHeader(500)
		if appMode == "development" {
			errorMsg = fmt.Sprintf("Failed to load template: %s\n", err.Error())
		} else {
			fmt.Errorf("Failed to load template: %s\n", err.Error())
			errorMsg = "Sorry! We hit an error. Please reload and try again"
		}
		fmt.Fprintf(w, errorMsg)
	} else {
		t.Execute(w, struct{ DevMode bool }{appMode == "development"})
	}
}

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handleHomePage(w, r)
		} else {
			// Serve assets
			fs.ServeHTTP(w, r)
		}
	})

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

	// LiveReload setup
	var done chan bool
	if appMode == "development" {
		lr, err := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
		if err != nil {
			fmt.Println("Unable to start livereload! %s\n", err.Error())
		}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("Unable to create file watcher: %s\n", err.Error())
		}
		defer watcher.Close()

		done = make(chan bool)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Write == fsnotify.Write {
						lr.Reload(event.Name)
					}
				case err := <-watcher.Errors:
					log.Println("Watcher error:", err)
				}
			}
		}()

		// Add dir to watcher
		err = watcher.Add("public/js")
		if err != nil {
			fmt.Println("Unable to set up watcher on public/js folder: %s\n", err.Error())
		}
		err = watcher.Add("public/css")
		if err != nil {
			fmt.Println("Unable to set up watcher on public/css folder: %s\n", err.Error())
		}

		fmt.Println("No errors so far? Livereload watching public folder")
		go lr.ListenAndServe()
		// Start goroutine that requests reload upon watcher event
		go func() {
			for {
				event := <-watcher.Events
				lr.Reload(event.Name)
			}
		}()
	}

	fmt.Printf("Serving files from %s on port %s\n", photosPath, port)
	http.ListenAndServe(":"+port, nil)
	<-done
	defer close(done)
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
