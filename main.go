// main.go
// Learned scraping from https://brightdata.com/blog/how-tos/web-scraping-go
// thousand ways to improve this:
// focus on: scraping with lib for maintainance, using argparse for further cli
// option/arg completion, flexibility in format output as a backend tool could
// be useful to parse it via http or cli for the different services around.

package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main() {
	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	collector.Visit("https://news.ycombinator.com/")
	fmt.Println("Hello, World!")
	// scraping logic...
}
