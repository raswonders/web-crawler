package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:];
	switch len(args) {
	case 0: 
		fmt.Println("no website provided")
		os.Exit(1)
	case 1: 
		url := args[0]
		pages := map[string]int{};
		fmt.Printf("starting crawl of: %s\n", url)
		crawlPage(url, url, pages)
		fmt.Println("--------------------------------------------------")
		fmt.Println("Pages visited:")
		for k, v := range pages {
			fmt.Printf("%s: %d\n", k, v)
		}
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

}