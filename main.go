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
		fmt.Printf("starting crawl of: %s\n", url)
		html, err := getHTML(url)
		if err != nil {
			fmt.Printf("Fetching html failed: %v\n", err)
		}
		fmt.Println(html)
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
}