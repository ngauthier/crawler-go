package main

import (
	"fmt"
	"github.com/ngauthier/crawler"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: crawl <url>\n")
		return
	}
	url := os.Args[1]
	crawler.Scrape(url)
}
