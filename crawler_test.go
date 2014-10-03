package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
)

var (
	mux     *http.ServeMux
	server  *httptest.Server
	crawler *Crawler
	n       *NanoTest
)

type NanoTest struct {
	t *testing.T
}

func (n *NanoTest) Die(err error) {
	if err != nil {
		n.t.Fatalf("Error: %v", err)
	}
}

func (n *NanoTest) Equal(a interface{}, b interface{}) {
	if a != b {
		trace := make([]byte, 1024)
		runtime.Stack(trace, false)
		n.t.Logf("%s", trace)
		n.t.Fatalf("Expected %v to equal %v", b, a)
	}
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func setup(t *testing.T) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	crawler = New(server.URL)
	n = &NanoTest{t: t}
}

func teardown() {
	server.Close()
}

func TestScrape(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `<a href='/'>Root</a>`)
	})

	// TODO: actually check the sitemap returned
	_, err := crawler.Scrape("/")
	n.Die(err)
}

func TestNewPage(t *testing.T) {
	setup(t)

	file, err := os.Open("test/fixtures/single_page/index.html")
	n.Die(err)

	doc, err := goquery.NewDocumentFromReader(file)
	n.Die(err)

	page := NewPage(doc)
	n.Equal(1, len(page.Links))
	n.Equal("/about", page.Links[0])

	// TODO: check css and scripts too
}

/*
 * TODO: scraping from a root page and having it scrape a referenced page into a sitemap struct
 * TODO: turning a page struct into json
 * TODO: turning a sitemap struct into json
 */
