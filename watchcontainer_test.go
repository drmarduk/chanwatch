package main

import (
	"log"
	"testing"
)

func TestNewWatchContainer(t *testing.T) {
	wc := NewWatchContainer()
	if wc.Lists == nil {
		t.Fatal("could not make WatchList map")
	}
}

func TestWatchContainerNewList(t *testing.T) {
	wc := NewWatchContainer()

	id := wc.NewList()
	wl, err := wc.GetWatchList(id)
	if err != nil {
		t.Fatalf("could not find added WatchList, %s\n", err)
	}

	if wl.ID != id {
		t.Fatal("id is not matching")
	}
}

func TestWatchContainerAddThread(t *testing.T) {
	url := "http://example.com"
	wc := NewWatchContainer()
	id := wc.NewList()

	err := wc.AddThread(id, url)
	if err != nil {
		t.Fatalf("could not add url to list, %v\n", err)
	}

	id = ""
	err = wc.AddThread(id, url)
	if err == nil {
		t.Fatal("should be err not found")
	}
}

func TestWatchContainerGet(t *testing.T) {
	url := "http://example.com"
	wc := NewWatchContainer()
	id := wc.NewList()

	err := wc.AddThread(id, url)
	if err != nil {
		log.Fatalf("could not add url to list, %v\n", err)
	}

	wl, err := wc.GetWatchList(id)
	if err != nil {
		t.Fatal("WatchList could not be found")
	}
	if wl.ID != id {
		t.Fatalf("WatchList ID mismatching")
	}
}
