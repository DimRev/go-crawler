package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseUrl, rawCurrentUrl string, pages map[string]int) (map[string]int, error) {
	// Get the HTML content of the current page
	htmlStr, err := getHTML(rawCurrentUrl)
	if err != nil {
		return nil, err
	}

	// Extract all the URLs from the HTML
	curUrlStrs, err := getURLsFromHTML(htmlStr, rawBaseUrl)
	if err != nil {
		return nil, err
	}

	// Parse the base URL once
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, err
	}

	// Iterate over extracted URLs
	for _, curUrlStr := range curUrlStrs {
		curUrl, err := url.Parse(curUrlStr)
		if err != nil {
			fmt.Println("Skipping invalid URL:", curUrlStr)
			continue // Skip this URL if parsing fails
		}

		// Only crawl pages from the same domain
		if _, ok := pages[curUrlStr]; !ok {
			pages[curUrlStr] = 1 // Initialize page count
			if baseUrl.Host == curUrl.Host {
				// Recursively crawl the new page
				_, err := crawlPage(rawBaseUrl, curUrlStr, pages)
				if err != nil {
					fmt.Println("Error crawling page:", curUrlStr)
					continue
				}
			}
		} else {
			// Increment page count if already visited
			pages[curUrlStr]++
		}
	}
	return pages, nil
}
