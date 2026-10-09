package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct{
		name string
		html string
		expected string
	}{
		{
			name: "finds <h1>",
			html: `<html>
	<body>
		<h1>Welcome to Boot.dev</h1>
		<h2>The best site</h2>
		<main>
			<p>Learn to code by building real projects.</p>
			<p>This is the second paragraph.</p>
		</main>
	</body>
</html>`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "falls back to <h2>",
			html: `<html>
	<body>
		<h2>The best site</h2>
		<main>
			<p>Learn to code by building real projects.</p>
			<p>This is the second paragraph.</p>
		</main>
	</body>
</html>`,
			expected: "The best site",
		},
		{
			name: "defaults to empty string",
			html: `<html>
	<body>
		<main>
			<p>Learn to code by building real projects.</p>
			<p>This is the second paragraph.</p>
		</main>
	</body>
</html>`,
			expected: "",
		},
	} 

	for i, tc := range tests {
		actual := getHeadingFromHTML(tc.html)
		if actual != tc.expected {
			t.Errorf("Test %v - '%s' FAIL: expected: %s actual: %v", i, tc.name, tc.expected, actual)
		}
	}

} 