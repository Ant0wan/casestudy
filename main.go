// main.go
// Learned scraping from https://brightdata.com/blog/how-tos/web-scraping-go
// then read gocolly colly documentation https://go-colly.org/docs/examples/basic/
// doc https://developer.mozilla.org/fr/docs/Web/API/Document/querySelector
// thousand ways to improve this:
// focus on: scraping with lib for maintainance, using argparse for further cli
// option/arg completion, flexibility in format output as a backend tool could
// be useful to parse it via http or cli for the different services around.
// paging to take into account
// 3 structs: top, athing, bottom

package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

type Link struct {
	Url       string
	Paths     []string
	Externals []string
}

func scrap(u *url.URL) Link {
	var paths []string
	var externals []string

	// Instantiate collector
	c := colly.NewCollector()

	// Scrapping logic
	c.OnHTML("a", func(e *colly.HTMLElement) {

		// retreive all href from all 'a'
		link := e.Attr("href")

		u, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}

		rel, err := u.Parse(link) // if -o output, else put in JSON
		if err != nil {
			log.Fatal(err)
		}

		if u.IsAbs() {
			externals = append(externals, rel.String())

		} else {
			paths = append(paths, rel.String())
		}

	})

	c.Visit(u.String())
	return Link{
		Url:       u.String(),
		Paths:     paths,
		Externals: externals,
	}
}

func main() {

	u, err := url.Parse("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}

	link := scrap(u)
	fmt.Println(link)

}
