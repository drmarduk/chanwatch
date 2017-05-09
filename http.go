package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// indexHandler collects all threads available from
// all chans and presents them on the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.New("index")
	tmp, err := tmp.ParseFiles("./html/index.html")
	if err != nil {
		log.Printf("error while parsing template. %v\n", err)
		return
	}

	mutex.Lock()
	t := threads
	mutex.Unlock()

	data := struct {
		URL []Thread
	}{
		URL: t,
	}

	tmp.Execute(w, &data)
}

// TODO: make this an own type
var watchContainer WatchContainer

func newwatchlistHandler(w http.ResponseWriter, r *http.Request) {
	id := watchContainer.NewList()

	url := fmt.Sprintf("/watchlist/view/%s", id)
	http.Redirect(w, r, url, 302)
}

func viewwatchlistHandler(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Path
	req = strings.Replace(req, "/watchlist/view/", "", -1)

	log.Printf("/viewwatchlist/view/%s .\n", req)

	wl, err := watchContainer.GetWatchList(req)
	if err != nil {
		// no list found
		log.Printf("WatchList %s not found!", req)
		http.Redirect(w, r, "/", 302)
		return
	}

	log.Printf("Found Watchlist: %s.\n", wl.String())

	tmp := template.New("view_watchlist")
	tmp, err = tmp.ParseFiles("./html/view_watchlist.html")
	if err != nil {
		log.Printf("error while parsing template. %v\n", err)
		return
	}

	data := struct {
		Base      string
		Watchlist *WatchList
	}{
		Base:      fmt.Sprintf("http://localhost:8080/watchlist/view/%s", wl.ID),
		Watchlist: wl,
	}

	tmp.Execute(w, &data)
}

func addwatchlistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	thread := r.FormValue("thread")
	log.Printf("/watchlist/add: ID: %s - Thread: %s\n", id, thread)
	// TODO: sanitize threadurl and id
	watchContainer.AddThread(id, thread)
	// TODO: do proper redirection
	url := fmt.Sprintf("/watchlist/view/%s", id)
	http.Redirect(w, r, url, 302)
}
