// cmd.go

// The "cmd" package serves as the command-line interface for the myprogram tool.
// It uses the Cobra library for building a user-friendly command-line application.
// The primary functionality includes defining command-line flags, parsing arguments,
// and triggering the execution of the web scraping process through the lib.Worker function.
// The lib.Worker function is expected to be defined in the "myprogram/lib" package,
// handling the actual web scraping logic for a given URL.
// Note: The lib.Worker function is assumed to be defined in the lib package.
// It is responsible for performing the actual scraping of a given URL.

package cmd

import (
	"log"
	"os"
	"sync"

	"myprogram/lib" // Importing the custom lib package

	"github.com/spf13/cobra" // Importing the Cobra library for building command-line applications
)

// Command-line flags and variables
var (
	addrs   []string // Slice to store URLs to be scrapped
	format  string   // Output format (json or stdout)
	rootCmd = &cobra.Command{
		Use:   "myprogram",
		Short: "A simple URL scrapper.",
		Long: `Given any number of HTTP URLs as command line parameters,
myprogram connects to each URL and extracts all links from it.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(addrs) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(1)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			scrap(addrs, format)
		},
	}
)

// scrap function takes a slice of URLs and an output format,
// spawns a worker for each URL, and waits for all workers to finish.
func scrap(addrs []string, format string) {
	var wg sync.WaitGroup

	for _, addr := range addrs {
		wg.Add(1)

		go func(addr string) {
			defer wg.Done()
			lib.Worker(addr, format)
		}(addr)
	}

	wg.Wait()
}

// Execute function starts the execution of the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init function initializes the command-line flags for the root command.
func init() {
	rootCmd.Flags().StringArrayVarP(&addrs, "url", "u", []string{}, "webpage URL to be scrapped")
	rootCmd.PersistentFlags().StringVarP(&format, "output", "o", "json", "output format: json or stdout (full URL)")
}
