package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct{
		name string
		url string 
		html string
		expected []string
	}{
		{
			name: "finds all urls",
			url: "https://crawler-test.com",
			html: `<html>
	<body>
		<a href="https://crawler-test.com">
			<span>Boot.dev</span>
		</a>
		<a href="https://crawler-test.com/something">
			<span>Boot.dev/something</span>
		</a>
	</body>
</html>`,
			expected: []string{"https://crawler-test.com",
			"https://crawler-test.com/something"},
		},
		{
			name: "finds relative urls",
			url: "https://crawler-test.com",
			html: `<html>
	<body>
		<a href="/something">
			<span>Boot.dev/something</span>
		</a>
	</body>
</html>`,
			expected: []string{"https://crawler-test.com/something"},
		},
		{
			name: "defaults to []",
			url: "https://crawler-test.com",
			html: `<html><body></body></html>`,
			expected: []string{},
		},
	}
	

	for i, tc := range tests {
		baseURL, err := url.Parse(tc.url)
		if err != nil {
			t.Errorf("couldn't parse input URL: %v", err)
			return
		}

		actual, err := getURLsFromHTML(tc.html, baseURL)
		if !reflect.DeepEqual(actual, tc.expected) || err != nil {
			t.Errorf("Test %v - '%s' FAIL: expected: %s actual: %v", i, tc.name, tc.expected, actual)
		}
	}
} 