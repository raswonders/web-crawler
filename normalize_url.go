package main

import "net/url"

func normalizeURL(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return u.Host+u.Path, nil
}