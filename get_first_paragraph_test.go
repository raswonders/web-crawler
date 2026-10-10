
package main

import "testing"

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct{
		name string
		html string
		expected string
	}{
		{
			name: "finds first <p>",
			html: `<html>
	<body>
		<h1>Welcome to Boot.dev</h1>
		<h2>The best site</h2>
		<p>Outside main and not so useful</p>
		<main>
			<p>Learn to code by building real projects.</p>
			<p>This is the second paragraph.</p>
		</main>
	</body>
</html>`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "falls back to <p> outside <main>",
			html: `<html>
	<body>
		<h1>Welcome to Boot.dev</h1>
		<h2>The best site</h2>
		<p>Outside main and not so useful</p>
		<main>
		</main>
	</body>
</html>`,
			expected: "Outside main and not so useful",
		},
		{
			name: "defaults to empty string",
			html: `<html>
	<body>
		<h1>Welcome to Boot.dev</h1>
		<h2>The best site</h2>
		<main>
		</main>
	</body>
</html>`,
			expected: "",
		},
	} 

	for i, tc := range tests {
		actual := getFirstParagraphFromHTML(tc.html)
		if actual != tc.expected {
			t.Errorf("Test %v - '%s' FAIL: expected: %s actual: %v", i, tc.name, tc.expected, actual)
		}
	}

} 