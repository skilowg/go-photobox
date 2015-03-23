package photobox

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	photosPath := os.Getenv("FILES")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths, err := List(photosPath)
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
