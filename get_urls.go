package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	res := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return res, err
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		if v, ok := s.Attr("href"); ok == true {
			url, err := url.Parse(v)
			if err != nil {
				fmt.Println("href couldn't be parsed as url") 
			}
			resolved := baseURL.ResolveReference(url)
			res = append(res, resolved.String())
		}
	})

	return res, nil
}
