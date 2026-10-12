package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:];
	switch len(args) {
	case 0: 
		fmt.Println("no website provided")
		os.Exit(1)
	case 1: 
		rawURL := args[0]
		currentURL, err := url.Parse(rawURL)
		if err != nil {
			fmt.Printf("couldn't parse url %v", err)
			return	
		}
		cfg := config{
			pages: map[string]PageData{},
			baseURL: &url.URL{
				Scheme: currentURL.Scheme,
				Host: currentURL.Host,
			},
			mu: &sync.Mutex{},
			concurrencyControl: make(chan struct{}, 1),
			wg: &sync.WaitGroup{},
		}
		cfg.wg.Add(1)
		cfg.crawlPage(rawURL)
		cfg.wg.Wait()
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
}