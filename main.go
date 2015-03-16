package main

import (
	"net/http"
	"os"

)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":"+port, nil)
}
