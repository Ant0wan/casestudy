// main.go
// argparse for further cli
// option/arg completion, flexibility in format output json, raw, yaml as a backend tool could
// be useful to parse it via http or cli for the different services around.
// paging to take into account !!
// add log level info, debug, fatal...

package main

import (
	"encoding/json"
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

func output(u *url.URL, link Link, oformat string) {
	switch oformat {
	case "json":
		jsonu := u.Scheme + "://" + u.Host
		output, err := json.Marshal(map[string][]string{jsonu: link.Paths})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	case "output":
		for _, path := range link.Paths {
			o, err := url.JoinPath(link.Url, path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(o)
		}
	default:
		log.Fatal("Format specified does not exist")
	}
}

func main() {

	u, err := url.Parse("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}

	link := scrap(u)

	output(u, link, "output")
}

// TODO: json, output formats
// Add option for all links not just relative
// Add option to sort output
// Use argparse
// add --all option
// Comment all code
// Loop over url args
// Add log level "cannot process URL"
// Multithread each URL in arg
