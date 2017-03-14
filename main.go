package main

import (
	"io"
	"net/http"
)

var threads []string

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ello")
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add")
}
