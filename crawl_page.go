package main

import (
	"fmt"
	"net/url"
	"time"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	fmt.Printf("crawling %s\n", rawCurrentURL)

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't parse current URL %q: %v\n", rawCurrentURL, err)
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("couldn't parse base URL %q: %v\n", rawBaseURL, err)
		return
	}

	if currentURL.Host != baseURL.Host {
		fmt.Printf("current url %q is not from same domain as base %q\n", currentURL, baseURL)
		return
	}	

	url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't normalize current URL %q: %v\n", currentURL, err)
		return 
	}

	_, isVisited := pages[url];
	if isVisited {
		pages[url] = pages[url] + 1
		return
	}
	pages[url] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't fetch html %v\n", err)
		return
	}
	fmt.Println(html)
	time.Sleep(500 * time.Millisecond)

	nextURLs, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("couldn't get urls: %v\n", err)
	}

	for _, u := range nextURLs {
		crawlPage(rawBaseURL, u, pages)
	}
}