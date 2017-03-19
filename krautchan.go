package main

import (
	"fmt"
	"html"
	"regexp"
)

// TODO: might push to own package

type krautchan struct {
}

// BaseURL returns the base url of the imageboard
func (k krautchan) BaseURL() string {
	return "http://krautchan.net"
}

// CatalogURL returns a convenvient way to list all threads
func (k krautchan) CatalogURL() string {
	return "http://krautchan.net/catalog/b"
}

func (k krautchan) ExtractThreads(src string) (result []Thread) {
	r := regexp.MustCompile(`class="thread_OP" id="(?P<id>[0-9]*)">\s*<div class="post">\s*<header>\s*<a href="(.*)"><h1>\s*(?P<title>(.*))`)

	names := r.SubexpNames()
	matches := r.FindAllStringSubmatch(src, -1)

	for _, m := range matches {
		md := map[string]string{}
		for i, n := range m {
			md[names[i]] = n
		}

		t := Thread{}
		t.URL = fmt.Sprintf("%s/b/thread-%s.html", k.BaseURL(), md["id"])
		t.Title = html.UnescapeString(md["title"])

		result = append(result, t)
	}
	return result
}
