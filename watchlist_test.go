package main

import (
	"testing"
	"time"
)

func TestNewWatchList(t *testing.T) {
	t1 := time.Now()
	wl := NewWatchList()
	if wl.Created.Sub(t1).Nanoseconds() < int64(0) {
		t.Fatalf("error while creating time stuff")
	}

	if len(wl.ID) != 5 {
		t.Fatalf("error while generating id, length not 5")
	}
}

func TestAddThread(t *testing.T) {
	url := "http://example.com"
	wl := NewWatchList()

	err := wl.AddThread(url)
	if err != nil {
		t.Fatalf("error while adding valid url")
	}
	found := false
	for _, t := range wl.Threads {
		if t.URL == url {
			found = true
		}
	}

	if !found {
		t.Fatal("could not find added thread in list")
	}

	err = wl.AddThread("")
	if err == nil {
		t.Fatal("err should not be nil")
	}
}
