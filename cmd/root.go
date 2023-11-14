package cmd

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goscrape",
	Short: "GoScrape is a CLI tool for efficient web scraping based on sitemaps.",
	Long: `
	GoScrape is designed to simplify the process of extracting data from websites. 
	It uses sitemaps to navigate through pages and can handle various output formats, 
	including Excel. With GoScrape, users can perform structured scraping operations via 
	simple CLI commands, configure scraping parameters, and manage output settings with ease. 
	Whether you're aggregating data for analysis or archiving web content, GoScrape provides 
	a powerful and user-friendly interface for all your web scraping needs.`,
}

func init() {
	rootCmd.SetHelpFunc(helpFunction)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(startCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func helpFunction(cmd *cobra.Command, args []string) {
	header := color.New(color.FgHiMagenta)
	command := color.New(color.FgHiCyan)
	option := color.New(color.FgHiGreen)
	link := color.New(color.FgBlue)
	example := color.New(color.FgWhite)

	helpMessage := "\n"
	helpMessage += fmt.Sprintf("GoScrape CLI %v\n\n", emoji.Sprint(":globe_with_meridians:"))
	helpMessage += header.Sprint("Usage:\n")
	helpMessage += "  " + command.Sprint("goscrape") + " [command] [arguments]\n\n"
	helpMessage += header.Sprint("Commands:\n")
	helpMessage += "  " + command.Sprint("start") + "     Initiate the web scraping process with a given sitemap URL.\n"
	helpMessage += "             Example:\n"
	helpMessage += "               " + example.Sprint("goscrape start https://example.com/sitemap.xml") + "\n\n"
	helpMessage += "  " + command.Sprint("config") + "    View or set configuration settings for the scraper.\n"
	helpMessage += "             Example:\n"
	helpMessage += "               " + example.Sprint("goscrape config --set output=json") + "\n\n"
	helpMessage += "  " + command.Sprint("status") + "    Check the current status of the scraper.\n"
	helpMessage += "             Example:\n"
	helpMessage += "               " + example.Sprint("goscrape status") + "\n\n"
	helpMessage += header.Sprint("Options:\n")
	helpMessage += "  " + option.Sprint("-h, --help") + "    Show this help message and exit.\n"
	helpMessage += "  " + option.Sprint("-v, --version") + " Show the version number and exit.\n\n"
	helpMessage += "For more information or to report issues, visit " + link.Sprint("https://github.com/ngfenglong/go-scrape-flow") + "\n"

	fmt.Println(helpMessage)
}
