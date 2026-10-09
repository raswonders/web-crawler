package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	res := doc.Find("h1")
	if len(res.Nodes) > 0 {
		return res.Text()
	}

	res = doc.Find("h2")
	if len(res.Nodes) > 0 {
		return res.Text()
	}

	return ""
}