// main.go

// Package main provides the main entry point for the program.
// The program is designed to crawl a given URL, extract internal and external links,
// and output the results in either JSON or stdout format.

// Note: Paging has not been implemented in this version.
// TODO: Add an option for all links, not just relative (consider adding a --all option)
// TODO: Add log level "cannot process URL"
// TODO: Add comprehensive comments to explain the code in more detail

package main

import (
	"myprogram/cmd" // Importing the custom cmd package
)

// main function is the entry point of the program.
// It executes the cmd package and starts the program.
func main() {
	cmd.Execute()
}
