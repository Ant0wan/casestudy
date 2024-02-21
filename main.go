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
	"sync"

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

		rel, err := u.Parse(link)
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

func worker(addr string, oformat string) {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}

	link := scrap(u)

	output(u, link, oformat)
}

func main() {

	var wg sync.WaitGroup
	addrs := []string{"https://news.ycombinator.com/", "https://arstechnica.com/"}

	for _, addr := range addrs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(addr, "output")
		}()
	}

	wg.Wait()
}

// Add option for all links not just relative add --all option
// Add option to sort output
// Use argparse

// Add log level "cannot process URL"

// Comment all code
