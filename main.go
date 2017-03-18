package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	timeout = 30
)

var mutex sync.Mutex
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

		tmp := board.ExtractThreads(src)

		// lock that shit and release afterwards
		mutex.Lock()
		threads = tmp
		mutex.Unlock()

		log.Printf("Found %d threads.\n", len(threads))
		for _, t := range threads {
			log.Println(t.URL + " - " + t.Title)
		}

		time.Sleep(timeout * time.Second)
	}
}
