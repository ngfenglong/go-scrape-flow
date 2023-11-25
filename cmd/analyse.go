package cmd

import (
	"time"

	"github.com/ngfenglong/go-scrape-flow/helper/excel"
	"github.com/ngfenglong/go-scrape-flow/sitemap"
	"github.com/spf13/cobra"
)

var analyseCmd = &cobra.Command{
	Use:   "analyse [URL]",
	Short: "Analyse a webpage and list common headers and example text.",
	Long: `
		The 'analyse' command is designed to help users identify potential selectors 
		for web scraping by analyzing a given webpage. When executed, it parses the 
		webpage's HTML content and extracts information about various HTML elements, 
		such as class names, IDs, and other attributes that can be used as selectors.

		This command provides a comprehensive overview of the webpage's structure, 
		highlighting common headers and providing example text or content for each 
		identified selector. This insight is particularly useful for users who wish 
		to perform targeted scraping but are unsure which selectors to use.

		The output of the 'analyse' command includes:

		- A list of unique selectors found on the webpage, such as class names and IDs.
		- Example content for each selector, which helps in understanding what kind of 
		data each selector will yield when used in scraping.
		- The frequency or count of each selector, giving an idea of how often it 
		appears on the page.

		Note that the 'analyse' command focuses on static HTML content. Therefore, 
		dynamically loaded content (such as that loaded via JavaScript) might not be 
		fully analysed. The command is best used for initial reconnaissance of a 
		webpage's structure and to aid in formulating a scraping strategy.

		This command is an essential tool for both novice and experienced users of 
		GoScrapeFlow, streamlining the process of setting up a scraping task by 
		providing clear and actionable information about potential data points on 
		a target webpage.`,
	Args: cobra.MinimumNArgs(1),
	Run:  runAnalyse,
}

func init() {
	analyseCmd.PersistentFlags().StringVarP(&outputFileName, "output", "o", "", "Optional: Name of the output file")
}

func runAnalyse(cmd *cobra.Command, arg []string) {
	url := arg[0]

	var fileName string
	if outputFileName != "" {
		fileName = outputFileName
	} else {
		dateStr := time.Now().Format("02012006")
		fileName = "Analysis_" + dateStr
	}

	content := sitemap.AnalysePageStructure(url)
	createAnaylsedDataExcelFile(content, fileName)

}

func createAnaylsedDataExcelFile(content sitemap.ContentMap, fn string) {
	f := excel.CreateNewFile(fn)
	defer f.Close()

	tagSheetName := "Tag"
	f.CreateNewSheet(tagSheetName)

	tagHeader := []interface{}{"Tag", "Example"}
	f.SetRowValues(tagSheetName, 1, tagHeader)

	r := 2
	for key, value := range content.Tags {
		row := []interface{}{key, value}
		f.SetRowValues(tagSheetName, r, row)
		r++
	}

	classSheetName := "Classes"
	f.CreateNewSheet(classSheetName)

	classHeader := []interface{}{"Class", "Example"}
	f.SetRowValues(classSheetName, 1, classHeader)

	r = 2
	for key, value := range content.Classes {
		row := []interface{}{key, value}
		f.SetRowValues(classSheetName, r, row)
		r++
	}

	idSheetName := "Ids"
	f.CreateNewSheet(idSheetName)

	idHeader := []interface{}{"IDs", "Example"}
	f.SetRowValues(idSheetName, 1, idHeader)

	r = 2
	for key, value := range content.Tags {
		row := []interface{}{key, value}
		f.SetRowValues(idSheetName, r, row)
		r++
	}

	f.SaveAs("./output/")
}
