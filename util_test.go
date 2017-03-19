package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func TestGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)
	server := httptest.NewServer(mux)
	defer server.Close()

	src, err := get(server.URL)
	if err != nil {
		t.Fatalf("error while downloading, %v\n", err)
	}

	if src != "hello world" {
		t.Fatal("content mismatch")
	}

	if src == "" && err != nil {

	}
}
