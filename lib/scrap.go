// lib.go

// lib.go provides functionalities for web scraping. It defines the 'crawl' function,
// which uses the colly package to extract internal and external links from a given URL.
// The 'output' function formats and prints the results based on the specified output format.
// Additionally, the 'Worker' function acts as a higher-level interface, combining the crawl and output logic,
// facilitating web scraping on a URL with a specified output format.

package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly" // Importing the colly package for web scraping
)

// Link represents a structure to store information about a URL,
// including its paths and external links.
type Link struct {
	Url       string   // The original URL
	Paths     []string // Internal paths found on the page
	Externals []string // External links found on the page
}

// crawl function takes a URL, performs web scraping, and returns a Link structure.
func crawl(u *url.URL) Link {
	var paths []string     // Internal paths found on the page
	var externals []string // External links found on the page

	// Instantiate collector for web scraping
	c := colly.NewCollector()

	// OnHTML is called when an HTML element matching the given selector is found.
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// Retrieve all href attributes from anchor tags
		link := e.Attr("href")

		u, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}

		rel, err := u.Parse(link)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the link is absolute or relative and store accordingly
		if u.IsAbs() {
			externals = append(externals, rel.String())

		} else {
			paths = append(paths, rel.String())
		}
	})

	// Visit the specified URL for web scraping
	c.Visit(u.String())

	// Return the Link structure with the extracted information
	return Link{
		Url:       u.String(),
		Paths:     paths,
		Externals: externals,
	}
}

// output function takes a URL, Link structure, and output format as input,
// and prints the results in the specified format (JSON or stdout).
func output(u *url.URL, link Link, oformat string) {
	switch oformat {
	case "json":
		// Creating a JSON representation of the internal paths
		jsonu := u.Scheme + "://" + u.Host
		output, err := json.Marshal(map[string][]string{jsonu: link.Paths})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))

	case "stdout":
		// Printing each internal path
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

// Worker function takes an address and output format,
// parses the URL, performs web scraping, and prints the results.
func Worker(addr string, oformat string) {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}

	link := crawl(u)

	output(u, link, oformat)
}
