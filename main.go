package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed index.html
var index []byte

//go:embed favicon.png
var favicon []byte

func main() {

	http.HandleFunc("/favicon", func(w http.ResponseWriter, r *http.Request) {
		w.Write(favicon)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(index)
	})

	http.HandleFunc("/title", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.FormValue("title"))
	})

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
