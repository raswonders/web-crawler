package main

import "net/url"

type PageData struct {
	URL string 
	Heading string
	FirstParagraph string
	OutgoingLinks []string
	ImageURLs []string
}

func extractPageData(html, pageURL string) PageData {
	var pageData PageData	
	pageData.URL = pageURL 
	pageData.Heading = getHeadingFromHTML(html)
	pageData.FirstParagraph = getFirstParagraphFromHTML(html)

	url, err := url.Parse(pageURL)
	if err != nil {
		return pageData 
	}
	pageData.OutgoingLinks, _ = getURLsFromHTML(html, url)
	pageData.ImageURLs, _ = getImagesFromHTML(html, url)

	return pageData
}
