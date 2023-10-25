package sitemap

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	httpclient "github.com/ngfenglong/web-scrapper/http"
)

func CrawlSitemapURLs(url string) ([]string, []string) {
	worklist := make(chan []string)
	toScrape := []string{}
	errLog := []string{}
	n := 1

	go func() { worklist <- []string{url} }()

	for ; n > 0; n-- {
		list := <-worklist
		// For each link, create a new goroutine to extract of sitemaps and pages
		for _, link := range list {
			n++ // to block and it is looping through all worklist
			go func(link string) {
				res, err := httpclient.GetRequest(link)
				if err != nil {
					errLog = append(errLog, fmt.Sprint("Error making request with: ", link))
				} else {
					urls, err := extractUrlsFromResponse(res)
					if err != nil {
						errLog = append(errLog, fmt.Sprint("Error extracting urls from: ", link))
					}
					sitesMap, pages := segregateURLs(urls)

					if sitesMap != nil {
						worklist <- sitesMap
					}

					toScrape = append(toScrape, pages...)
				}
			}(link)
		}
	}

	return toScrape, errLog
}

func extractUrlsFromResponse(res *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []string{}, err
	}

	results := []string{}
	locSelectors := doc.Find("loc")
	for i := range locSelectors.Nodes {
		loc := locSelectors.Eq(i)
		result := loc.Text()
		results = append(results, result)
	}

	return results, nil
}

func segregateURLs(urls []string) ([]string, []string) {
	siteMaps := []string{}
	pages := []string{}

	for _, url := range urls {
		isSiteMap := strings.Contains(url, ".xml")
		if isSiteMap {
			siteMaps = append(siteMaps, url)
		} else {
			pages = append(pages, url)
		}
	}

	return siteMaps, pages
}
