package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thngkaiyuan/look-at-my-site/api"
	"github.com/abiosoft/semaphore"
)

const (
	port = "8080"
)

func main() {
	staticHandler := http.FileServer(http.Dir("web"))
	http.Handle("/", staticHandler)
	http.Handle("/favicon.ico", staticHandler)
	http.Handle("/static/", staticHandler)

	s := semaphore.New(3)
	a := api.New(s)
	http.HandleFunc("/api/check", a.Check)

	fmt.Printf("Start listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
