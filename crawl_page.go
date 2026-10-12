package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	fmt.Printf("crawling %s\n", rawCurrentURL)
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't parse current URL %q: %v\n", rawCurrentURL, err)
		return
	}


	if currentURL.Host != cfg.baseURL.Host {
		fmt.Printf("current url %q is not from same domain as base %q\n", currentURL, cfg.baseURL)
		return
	}	

	url, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't normalize current URL %q: %v\n", currentURL, err)
		return 
	}

	cfg.mu.Lock()
	_, isVisited := cfg.pages[url];
	if isVisited {
		fmt.Println("visited! leaving")
		cfg.mu.Unlock()
		return
	}
	cfg.pages[url] = PageData{}
	cfg.mu.Unlock()

	cfg.concurrencyControl <- struct{}{};
	html, err := getHTML(rawCurrentURL)
	<-cfg.concurrencyControl
	if err != nil {
		fmt.Printf("couldn't fetch html %v\n", err)
		return
	}
	fmt.Println(html)

	page := extractPageData(html, rawCurrentURL);
	cfg.mu.Lock()
	cfg.pages[url] = page 
	cfg.mu.Unlock()

	for _, u := range page.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}