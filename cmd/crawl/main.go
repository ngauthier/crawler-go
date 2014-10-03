package main

import (
	"fmt"
	"github.com/ngauthier/crawler"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: crawl <url>\n")
		return
	}
	url, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Error parsing url: %v\n", err)
	}
	// TODO: split the argument better so we retain query params and fragments in starting path
	// Possibly change the crawler API to accept a full url, and split up internally for external
	// link checking
	c := crawler.New(url.Scheme + "://" + url.Host)
	c.Scrape(url.Path)
}
