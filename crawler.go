package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Crawler struct {
	Host string
}

type Page struct {
	Links []string
}

func NewCrawler(host string) *Crawler {
	crawler := &Crawler{Host: host}
	return crawler
}

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

func NewPage(doc *goquery.Document) *Page {
	p := &Page{}

	p.Links = links(doc)

	return p
}
