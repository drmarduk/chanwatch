package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"
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

func krautchanExtractTitle(src string) string {
	r := regexp.MustCompile(`<span class="postsubject">(.*)</span>`)
	match := r.FindString(src)
	match = match[26:]
	match = strings.Replace(match, "</span>", "", 1)
	if match == "" {
		return "no title found"
	}
	return match
}

func krautchanExtractImages(src string) []string {
	r := regexp.MustCompile(`<a href="/files/(.*).(jpg|jpeg|gif|png|webm)" target="_blank">`)
	matches := r.FindAllString(src, -1)

	var result []string

	for _, m := range matches {
		m = m[9:]
		m = strings.Replace(m, `" target="_blank">`, "", 1)
		m = "http://krautchan.net" + m

		result = append(result, m)
	}
	return result
}
