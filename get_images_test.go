package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct{
		name string
		url string 
		html string
		expected []string
	}{
		{
			name: "finds all absolute image urls",
			url: "https://crawler-test.com",
			html: `<html>
	<body>
		<img src="https://crawler-test.com/image1.png" alt="Image">
		<img src="https://crawler-test.com/image2.png" alt="Image">
	</body>
</html>`,
			expected: []string{"https://crawler-test.com/image1.png",
			"https://crawler-test.com/image2.png"},
		},
		{
			name: "finds all relative image urls",
			url: "https://crawler-test.com",
			html: `<html>
	<body>
		<img src="/image1.png" alt="Image">
		<img src="/image2.png" alt="Image">
	</body>
</html>`,
			expected: []string{"https://crawler-test.com/image1.png",
			"https://crawler-test.com/image2.png"},
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

		actual, err := getImagesFromHTML(tc.html, baseURL)
		if !reflect.DeepEqual(actual, tc.expected) || err != nil {
			t.Errorf("Test %v - '%s' FAIL: expected: %s actual: %v", i, tc.name, tc.expected, actual)
		}
	}
} 