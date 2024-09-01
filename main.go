package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]
	rawMaxConcurrency := os.Args[2]
	rawMaxPages := os.Args[3]

	maxConcurrency, err := strconv.Atoi(rawMaxConcurrency)
	if err != nil {
		fmt.Printf("Error - failed to strconv: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(rawMaxPages)
	if err != nil {
		fmt.Printf("Error - failed to strconv: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg)
}
