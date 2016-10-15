package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = "8080"
)

func main() {
	staticHandler := http.FileServer(http.Dir("web"))
	http.Handle("/", staticHandler)
	http.Handle("/favicon.ico", staticHandler)
	http.Handle("/static/", staticHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
