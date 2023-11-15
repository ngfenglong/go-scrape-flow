package cmd

import (
	"fmt"
	"time"

	"github.com/ngfenglong/go-scrape-flow/helper/excel"
	"github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/ngfenglong/go-scrape-flow/sitemap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logger = log.New(logrus.InfoLevel, false)
var outputFileName string

var startCmd = &cobra.Command{
	Use:   "start [sitemap URL]",
	Short: "Start the web scraping process with a given sitemap URL.",
	Long:  "Initiates web scarping based on the provided sitemap URL.",
	Args:  cobra.MinimumNArgs(1),
	Run:   runStart,
}

func init() {
	startCmd.PersistentFlags().StringVarP(&outputFileName, "output", "o", "", "Optional: Name of the output file")
}

func runStart(cmd *cobra.Command, args []string) {
	url := args[0]

	fmt.Println("Starting scrape for: ", url)

	data := scrapeSitemap(url, 10)

	var fileName string
	if outputFileName != "" {
		fileName = outputFileName
	} else {
		dateStr := time.Now().Format("02012006")
		fileName = "Excel_File_" + dateStr
	}

	createExcelFile(data, fileName)
}

func scrapeSitemap(url string, concurr int) []sitemap.Data {
	pages := sitemap.CrawlSitemapURLs(url)
	logger.Info("Crawling Pages - ", logrus.Fields{"Number of Page": len(pages)})
	data := sitemap.CrawlPages(pages, concurr)

	return data

}

func createExcelFile(data []sitemap.Data, fn string) {
	f := excel.CreateNewFile(fn)
	defer f.Close()

	var sheetName string = "Sheet1"
	f.CreateNewSheet(sheetName)
	f.SetRowValues(sheetName, 1, []interface{}{
		"URL",
		"Title",
		"MetaDescription",
		"H1",
	})

	for i, d := range data {
		f.SetRowValues(sheetName, i+2, []interface{}{
			d.URL,
			d.Title,
			d.MetaDescription,
			d.H1,
		})
	}

	f.SaveAs("./output/")
}
