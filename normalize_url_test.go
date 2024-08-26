package main

import (
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name        string
		inputUrl    string
		expected    string
		expectedErr string
	}{
		{name: "remove scheme", inputUrl: "https://blog.boot.dev/path", expected: "blog.boot.dev/path"},
		{name: "http scheme", inputUrl: "http://blog.boot.dev/path", expected: "blog.boot.dev/path"},
		{name: "invalid url", inputUrl: "invalid", expectedErr: "invalid URL: invalid"},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputUrl)
			if err != nil && tc.expectedErr == "" {
				if err.Error() != tc.expectedErr {
					t.Errorf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
					return
				}
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected %s, got %s", i, tc.name, tc.expected, actual)
			}
		})
	}
}
