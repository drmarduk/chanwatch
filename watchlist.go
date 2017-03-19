package main

import (
	"errors"
	"math/rand"
	"time"
)

// WatchList creates an in-memory storage for saving his own links
type WatchList struct {
	ID      string
	Threads []Thread

	Created time.Time
}

// NewWatchList asdf
func NewWatchList() WatchList {
	return WatchList{Created: time.Now(), ID: generateID(5)}
}

// AddThread appends a new Thread to the list
func (wl *WatchList) AddThread(url string) error {
	if url == "" {
		return errors.New("url is empty")
	}
	t := Thread{URL: url}

	// kA ob das alles überlebt
	go completeThreadInfo(t)
	wl.Threads = append(wl.Threads, t)
	return nil
}

func completeThreadInfo(t Thread) {
	// panic("complete shit")
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
