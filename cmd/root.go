package cmd

import (
	"os"
	"sync"

	"myprogram/lib"

	"github.com/spf13/cobra"
)

var (
	addr    string
	format  string
	rootCmd = &cobra.Command{
		Use:   "myprogram",
		Short: "A simple URL scrapper.",
		Long: `Given any number of HTTP URLs as command line parameters,
myprogram connects to each URL and extract all links from it.`,
		Run: func(cmd *cobra.Command, args []string) {
			var addrs []string
			addrs = append(addrs, addr)
			scrap(addrs)
		}}
)

func scrap(addrs []string) {
	var wg sync.WaitGroup

	for _, addr := range addrs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			lib.Worker(addr, "stdout")
		}()
	}

	wg.Wait()

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myprogram.yaml)")
	rootCmd.PersistentFlags().StringVarP(&addr, "url", "u", "", "webpage url to be scrapped")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.PersistentFlags().StringP("output", "FORMAT", "", "output either json or stdout(full url)")
	//	rootCmd.PersistentFlags().StringP("url", "URL", "", "url to be scrapped")
}
