package main

import (
	"testing"
)

func TestGetUrlsFromHtml(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		expectedErr string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<div>
				<span>Boot.dev</span>
			</div>
		</body>
	</html>
	`,
			expected: []string{},
		},
		{
			name:     "only relative paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="/path/two">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "only absolute paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/two">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://other.com/path/one", "https://other.com/path/two"},
		},
		{
			name:        "Invalid HTML",
			inputURL:    "https://blog.boot.dev",
			inputBody:   `This is just some text`,
			expectedErr: "couldn't parse HTML: html: parse error",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && tc.expectedErr == "" {
				if err.Error() != tc.expectedErr {
					t.Errorf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
					return
				}
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: unexpected length: expected %v, got %v", i, tc.name, tc.expected, actual)
				return
			}
			for j, actualURL := range actual {
				if actualURL != tc.expected[j] {
					t.Errorf("Test %v - %s FAIL: non-matching values in URL array, expected URLstring[%v]: %s, got %s",
						i,
						tc.name,
						j,
						tc.expected[j],
						actualURL,
					)
					return
				}
			}
		})
	}
}
