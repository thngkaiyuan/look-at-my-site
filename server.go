package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thngkaiyuan/look-at-my-site/api"
)

const (
	port = "8080"
)

func main() {
	staticHandler := http.FileServer(http.Dir("web"))
	http.Handle("/", staticHandler)
	http.Handle("/favicon.ico", staticHandler)
	http.Handle("/static/", staticHandler)

	a := api.New()
	http.HandleFunc("/api/check", a.Check)

	fmt.Printf("Start listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
