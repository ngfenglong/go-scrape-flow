package main

import (
	"fmt"

	"github.com/ngfenglong/web-scrapper/sitemap"
)

func main() {
	// Provided a sitemap URL
	url := "https://dummy.com/sitemap_index.xml"

	scrapeSitemap(url, 10)
}

func scrapeSitemap(url string, concurr int) {
	pages, _ := sitemap.CrawlSitemapURLs(url)

	data := sitemap.CrawlPages(pages, concurr)

	for _, d := range data {
		fmt.Print(d.H1)
	}
}
