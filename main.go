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
	"strings"

	"github.com/gocolly/colly"
)

type Link struct {
	Url   string
	Paths []string
}

func main() {
	// Instantiate default collector
	// Not sure configuration is good here
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// use a lib to filter properly all kind of url
		if strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://") {
			// use a lib to filte     r, _ := regexp.Compile("p([a-z]+)ch") string:// {
			fmt.Println("with protocol: ", link)
		} else { // add relative, remove the mailto: ... etc
			fmt.Println("relative: ", link)
		}

	})

	c.Visit("https://news.ycombinator.com/")
}
