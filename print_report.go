package main

import (
	"fmt"
	"sort"
)

func printReport(cfg *config) {
	// Step 1: Extract the keys (URLs) into a slice
	keys := make([]string, 0, len(cfg.pages))
	for url := range cfg.pages {
		keys = append(keys, url)
	}

	// Step 2: Sort the slice based on the map values (counts)
	sort.Slice(keys, func(i, j int) bool {
		return cfg.pages[keys[i]] > cfg.pages[keys[j]] // Sort in descending order
	})

	fmt.Printf("=============================\nREPORT for %s\n=============================\n", cfg.baseURL)
	for _, val := range keys {
		fmt.Printf("Found %d internal links to %s\n", cfg.pages[val], val)
	}
}
