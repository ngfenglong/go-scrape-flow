package main

import (
	"time"

	"github.com/ngfenglong/go-scrape-flow/helper/excel"
	log "github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/ngfenglong/go-scrape-flow/sitemap"
	"github.com/sirupsen/logrus"
)

var logger = log.New(logrus.InfoLevel, false)

func main() {
	// Provided a sitemap URL
	url := "https://www.rocktherankings.com/sitemap.xml"
	scrapeSitemap(url, 10)
}

func scrapeSitemap(url string, concurr int) {
	pages := sitemap.CrawlSitemapURLs(url)
	logger.Info("Crawling Pages - ", logrus.Fields{"Number of Page": len(pages)}) // Temporary
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
