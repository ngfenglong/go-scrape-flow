package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Configure settings for the web scraper.",
	Long:  "View or set configuration options for the web scraper. If no value is provided, the current setting for the key will be displayed.",
	Args:  cobra.MinimumNArgs(1), // Requires at least one argument: the key
	Run:   runConfig,
}

func runConfig(cmd *cobra.Command, args []string) {
	fmt.Print("config")
}
