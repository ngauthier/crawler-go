/*
Crawl a site and build a sitemap
*/
package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// Crawler represents a web crawler for a particular host (domain)
type Crawler struct {
	Host string
}

// Page is a single page and it extracts links, css, and scripts
type Page struct {
	Links []string
}

// Constructor to make a Crawler on the given host
func New(host string) *Crawler {
	crawler := &Crawler{Host: host}
	return crawler
}

// Scrape a site starting from the given path
func (c *Crawler) Scrape(path string) (string, error) {
	fmt.Printf("Scraping %s%s\n", c.Host, path)

	doc, err := goquery.NewDocument(c.Host + path)

	if err != nil {
		fmt.Printf("%v", err)
		panic("Can't fetch " + path)
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

	return "yay", nil
}

// Extract script sources from a goquery document
func scripts(doc *goquery.Document) []string {
	return query(doc, "script[src]", "src")

}

// Extract style sources from a goquery document
func styles(doc *goquery.Document) []string {
	return query(doc, "link[type=\"text/css\"]", "href")
}

// Extract links from a goquery document
func links(doc *goquery.Document) []string {
	return query(doc, "a", "href")
}

// Helper to extract an attribute from all matching tags
func query(doc *goquery.Document, query string, attribute string) []string {
	return doc.Find(query).Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr(attribute)
		return src
	})
}

// Constructor to make a new Page by scraping a goquery Document
func NewPage(doc *goquery.Document) *Page {
	p := &Page{}

	p.Links = links(doc)

	return p
}
