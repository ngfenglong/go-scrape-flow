package main

import (
	"os"
	"time"

	"github.com/ngfenglong/go-scrape-flow/cmd"
	"github.com/ngfenglong/go-scrape-flow/helper/excel"
	"github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/ngfenglong/go-scrape-flow/sitemap"
	"github.com/sirupsen/logrus"
)

var logger = log.New(logrus.InfoLevel, false)

func main() {
	// Provided a sitemap URL
	if err := cmd.Execute(); err != nil {
		logger.Error("Error executing command", logrus.Fields{"error": err})
		os.Exit(1)
	}

	// url := "https://www.rocktherankings.com/sitemap.xml"
	// ScrapeSitemap(url, 10)
}

func ScrapeSitemap(url string, concurr int) {
	pages := sitemap.CrawlSitemapURLs(url)
	logger.Info("Crawling Pages - ", logrus.Fields{"Number of Page": len(pages)})
	data := sitemap.CrawlPages(pages, concurr)

	dateStr := time.Now().Format("02012006")
	fileName := "Excel_File_" + dateStr

	createExcelFile(data, fileName)
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
