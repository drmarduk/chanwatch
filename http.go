package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

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

	wl, err := watchContainer.GetWatchList(req)
	if err != nil {
		// no list found
		log.Printf("WatchList %s not found!", req)
		http.Redirect(w, r, "/", 302)
	}

	log.Printf("Open WatchList %s\n", wl.ID)

	tmp := template.New("view_watchlist")
	tmp, err = tmp.ParseFiles("./html/view_watchlist.html")
	if err != nil {
		log.Printf("error while parsing template. %v\n", err)
		return
	}

	data := struct {
		Base      string
		Watchlist WatchList
	}{
		Base:      fmt.Sprintf("http://localhost:8080/watchlist/view/%s", wl.ID),
		Watchlist: wl,
	}

	tmp.Execute(w, &data)
}

func addwatchlistHandler(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Path
	req = strings.Replace(req, "/watchlist/add/", "", -1)

	thread := r.Form.Get("thread")
	// TODO: sanitize url
	watchContainer.AddThread(req, thread)
	// TODO: do proper redirection
	url := fmt.Sprintf("/watchlist/view/%s", req)
	http.Redirect(w, r, url, 302)
}
