package main

import "errors"

// WatchContainer organizes all threadlistsj
type WatchContainer struct {
	Lists map[string]WatchList
}

// NewWatchContainer returns a new instance of the container
func NewWatchContainer() WatchContainer {
	return WatchContainer{Lists: make(map[string]WatchList)}
}

// NewList creates a new watchlist and adds it to the container
func (wc *WatchContainer) NewList() string {
	wl := NewWatchList()
	wc.Lists[wl.ID] = wl
	return wl.ID
}

// AddThread adds a new thread to a sublist
func (wc *WatchContainer) AddThread(id, url string) error {
	wl, err := getList(id, wc.Lists)
	if err != nil {
		return err
	}
	wl.AddThread(url)
	return nil
}

// GetWatchList returns a List based on the id
func (wc *WatchContainer) GetWatchList(id string) (WatchList, error) {
	return getList(id, wc.Lists)
}

func getList(id string, lists map[string]WatchList) (WatchList, error) {
	var wl WatchList
	var ok bool

	if wl, ok = lists[id]; !ok {
		return wl, errors.New("watchlist not found in container")
	}
	return wl, nil
}
