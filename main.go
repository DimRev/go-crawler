package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProg[0]
	fmt.Println("starting crawler of:", baseUrl)
	pages := make(map[string]int)
	pages, err := crawlPage(baseUrl, baseUrl, pages)
	if err != nil {
		fmt.Println("error crawling page:", err)
		os.Exit(1)
	}

	// Create a slice to hold the map keys and values
	type kv struct {
		Key   string
		Value int
	}
	var sortedPages []kv
	for k, v := range pages {
		sortedPages = append(sortedPages, kv{k, v})
	}

	// Sort the slice based on the counts
	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Value > sortedPages[j].Value
	})

	// Print the sorted results
	for _, page := range sortedPages {
		fmt.Println(page.Key, page.Value)
	}
}
