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

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	// Not sure configuration is good here
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	c.OnHTML(".athing", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		// Print link
		fmt.Println(id)
	})

	c.Visit("https://news.ycombinator.com/")
}
