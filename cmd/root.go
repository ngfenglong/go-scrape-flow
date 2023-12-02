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
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(analyseCmd)
	rootCmd.SetHelpFunc(helpFunction)
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
	helpMessage += "             Options:\n"
	helpMessage += "               " + option.Sprint("--selector, -s") + "    Specify CSS selectors to scrape additional data.\n"
	helpMessage += "               " + option.Sprint("--output, -o") + "    Specify the output file name for the scraped data.\n"
	helpMessage += "             Example:\n"
	helpMessage += "               " + example.Sprint("goscrape start https://example.com/sitemap.xml -s '.class1' -s '#id1' -o output.xlsx") + "\n\n"

	helpMessage += "  " + command.Sprint("analyse") + "    Analyse a webpage and list common headers and example text.\n"
	helpMessage += "             Options:\n"
	helpMessage += "               " + option.Sprint("--output, -o") + "    Specify the output file name for the analysis results.\n"
	helpMessage += "             Example:\n"
	helpMessage += "               " + example.Sprint("goscrape analyse https://example.com -o analysis.xlsx") + "\n\n"

	helpMessage += header.Sprint("Options:\n")
	helpMessage += "  " + option.Sprint("-h, --help") + "    Show this help message and exit.\n"
	helpMessage += "  " + option.Sprint("-v, --version") + " Show the version number and exit.\n\n"
	helpMessage += "For more information or to report issues, visit " + link.Sprint("https://github.com/ngfenglong/go-scrape-flow") + "\n"

	fmt.Println(helpMessage)
}
