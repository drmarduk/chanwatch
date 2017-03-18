package main

// Imageboard is an interface to provide basic usage for many imageboards
type Imageboard interface {
	BaseURL() string
	CatalogURL() string
	ExtractThreads(string) []Thread
}

// Thread holds the basic metadata of a thread
type Thread struct {
	URL   string
	Title string
	Src   string // is optional
}
