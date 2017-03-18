package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.New("index")
	tmp, err := tmp.ParseFiles("./html/index.html")
	if err != nil {
		log.Printf("error while parsing template. %v\n", err)
		return
	}

	/*data := struct {
		URL []string
	}{
		URL: extractURLs(threads),
	}

	tmp.Execute(w, &data) */
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add")
}
