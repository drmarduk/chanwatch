package main

import (
	"log"
	"net/http"
	"time"
)

const (
	timeout = 30
)

var threads []Thread

func main() {

	kraut := krautchan{}
	go watcher(kraut)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func watcher(board Imageboard) {
	for {
		src, err := get(board.CatalogURL())
		if err != nil {
			log.Printf("error while downloading page: %v\n", err)
		}

		threads = board.ExtractThreads(src)

		log.Printf("Found %d threads.\n", len(threads))
		for _, t := range threads {
			log.Println(t.URL + " - " + t.Title)
		}

		time.Sleep(timeout * time.Second)
	}
}

func extractURLs(threads map[string]string) []string {
	var result []string
	for u, _ := range threads {
		result = append(result, u)
	}
	return result
}
