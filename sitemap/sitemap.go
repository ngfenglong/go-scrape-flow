package sitemap

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ngfenglong/go-scrape-flow/httpclient"
	"github.com/sirupsen/logrus"
)

func CrawlSitemapURLs(url string) []string {
	worklist := make(chan []string)
	toScrape := []string{}
	var n int
	n++
	go func() { worklist <- []string{url} }()

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			n++
			go func(link string) {
				res, err := httpclient.GetRequest(link)
				if err != nil {
					logger.Error("Error crawling sitemap ", logrus.Fields{"Error": err})
				} else {
					urls, err := extractUrlsFromResponse(res)
					if err != nil {
						logger.Error("Error extracting urls ", logrus.Fields{"Url": link, "Error": err})
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

	return toScrape
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
