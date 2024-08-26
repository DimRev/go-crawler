package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func getHTML(website string) (string, error) {
	currentURL, err := url.Parse(website)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	if currentURL.Scheme != "http" && currentURL.Scheme != "https" {
		return "", fmt.Errorf("invalid URL: %s", website)
	}

	resp, err := http.Get(website)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return "", fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(body), nil
}
