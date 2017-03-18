package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

var threads []string

func main() {

	go watcher()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":8080", nil)
}

func watcher() {
	for {
		src, err := get("http://krautchan.net/catalog/b")
		if err != nil {
			log.Printf("error while downloading page: %v\n", err)
		}

		threads = extractThreads(src)

		for _, t := range threads {
			log.Println(t)
		}

		time.Sleep(30 * time.Second)
	}
}

func extractThreads(src string) []string {
	var result []string

	r := regexp.MustCompile(`class="thread_OP" id="(?P<id>[0-9]*)">\s*<div class="post">\s*<header>\s*<a href="(.*)"><h1>(.*)<`) //(?P<title>(.*))<`)

	names := r.SubexpNames()
	matches := r.FindAllStringSubmatch(src, -1)

	for _, m := range matches {
		md := map[string]string{}
		for i, n := range m {
			md[names[i]] = n
		}

		result = append(result, fmt.Sprintf("http://krautchan.net/b/thread-%s.html - %s", md["id"], md["title"]))
	}
	return result
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.New("index")
	tmp, err := tmp.ParseFiles("./html/index.html")
	if err != nil {
		log.Printf("error while parsing template. %v\n", err)
		return
	}

	data := struct {
		URL []string
	}{
		URL: threads,
	}
	tmp.Execute(w, &data)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add")
}
