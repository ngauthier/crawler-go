package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: scrape <url>\n")
		return
	}
	url := os.Args[1]
	scrape(url)
}

func scrape(url string) {
	fmt.Printf("Scraping %s\n", url)

	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic("Can't fetch " + url)
	}

	fmt.Printf("Scripts:\n")
	for _, script := range scripts(doc) {
		fmt.Printf("\t%s\n", script)
	}

	fmt.Printf("Stylesheets:\n")
	for _, style := range styles(doc) {
		fmt.Printf("\t%s\n", style)
	}

	fmt.Printf("Links:\n")
	for _, link := range links(doc) {
		fmt.Printf("\t%s\n", link)
	}
}

func scripts(doc *goquery.Document) []string {
	return query(doc, "script[src]", "src")

}

func styles(doc *goquery.Document) []string {
	return query(doc, "link[type=\"text/css\"]", "href")
}

func links(doc *goquery.Document) []string {
	return query(doc, "a", "href")
}

func query(doc *goquery.Document, query string, attribute string) []string {
	return doc.Find(query).Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr(attribute)
		return src
	})
}
