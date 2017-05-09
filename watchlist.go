package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// WatchList creates an in-memory storage for saving his own links
type WatchList struct {
	ID      string
	Threads []*Thread

	Created time.Time
}

// NewWatchList asdf
func NewWatchList() *WatchList {
	return &WatchList{Created: time.Now(), ID: generateID(5)}
}

// String for some stats
func (wl *WatchList) String() string {
	return fmt.Sprintf("ID: %s - Links: %d - Created: %s", wl.ID, len(wl.Threads), wl.Created.Format(time.RFC822))
}

// AddThread appends a new Thread to the list
func (wl *WatchList) AddThread(url string) error {
	if url == "" {
		return errors.New("url is empty")
	}
	t := &Thread{URL: url}

	// kA ob das alles überlebt
	go completeThreadInfo(t)
	wl.Threads = append(wl.Threads, t)
	return nil
}

func completeThreadInfo(t *Thread) {
	// panic("complete shit")
	src, err := get(t.URL)
	if err != nil {
		log.Printf("could not get sourcecode %s: %v\n", t.URL, err)
		return
	}
	t.Src = src
	t.Title = krautchanExtractTitle(t.Src)
	t.Images = krautchanExtractImages(t.Src)
	t.ImageCount = len(t.Images)
}

func generateID(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
