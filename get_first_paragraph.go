package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// look in main selection first
	mainSelection := doc.Find("main")
	if len(mainSelection.Nodes) > 0 {
		res := mainSelection.Find("p")
		if len(res.Nodes) > 0 {
			return res.First().Text()
		} 
	}

	// when not found look in whole doc
	res := doc.Find("p")
	if len(res.Nodes) > 0 {
		return res.First().Text()
	} 

	return ""
}