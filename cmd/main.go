package main

import (
	"time"

	"github.com/ngfenglong/go-scrape-flow/helper/excel"
	"github.com/ngfenglong/go-scrape-flow/sitemap"
)

func main() {
	// Provided a sitemap URL
	url := "https://eatbook.sg/sitemap_index.xml"
	// url := "https://eatbook.sg/page-sitemap.xml"
	scrapeSitemap(url, 10)
}

func scrapeSitemap(url string, concurr int) {
	pages, _ := sitemap.CrawlSitemapURLs(url)
	data := sitemap.CrawlPages(pages, concurr)

	createExcelFile(data)
}

func createExcelFile(data []sitemap.Data) {
	dateStr := time.Now().Format("02012006")
	f := excel.CreateNewFile("Excel_File_" + dateStr)

	defer f.Close()

	// Add Header
	var sheetName string = "Sheet 1"
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
