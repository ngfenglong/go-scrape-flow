package sitemap

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/ngfenglong/go-scrape-flow/httpclient"
	"github.com/sirupsen/logrus"
)

type Data struct {
	URL             string
	Title           string
	MetaDescription string
	H1              string
	StatusCode      int
	ExtraData       map[string]string
}

var logger = log.New(logrus.InfoLevel, false)

var selectorsToRemove = []string{
	"body script",
	"body style",
	"body link",
}

func CrawlPages(pages []string, concurrency int, selectors []string) []Data {
	var wg sync.WaitGroup
	tokens := make(chan struct{}, concurrency) // To be use as Semaphore
	dataChan := make(chan Data, len(pages))    // Using channel so dont have to use mutex

	for _, link := range pages {
		if link != "" {
			logger.Info("Scraping Page ", logrus.Fields{"Url": link})
			wg.Add(1)
			go func(pageUrl string, tokens chan struct{}) {
				defer wg.Done()
				tokens <- struct{}{}
				logger.Info("Requesting from URL ", logrus.Fields{"Url": pageUrl})
				<-tokens
				res, err := httpclient.GetRequest(pageUrl)
				if err != nil {
					logger.Error("Error crawling page", logrus.Fields{"Url": pageUrl, "Error": err})
					return
				}

				data, err := extractContentFromResponse(res, selectors)
				if err != nil {
					logger.Error("Error extracting content from response", logrus.Fields{"Error": err})
					return
				}
				dataChan <- data
			}(link, tokens)
		}
	}

	go func() {
		wg.Wait()
		close(dataChan)
	}()

	var finalResults []Data
	for data := range dataChan {
		finalResults = append(finalResults, data)
	}

	return finalResults
}

func extractContentFromResponse(res *http.Response, selectors []string) (Data, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Data{}, err
	}

	// To remove unnecessary selectors if printing of the HTML is required
	pruneDocument(doc, selectorsToRemove)

	result := Data{}
	result.URL = res.Request.URL.String()
	result.StatusCode = res.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]").Attr("content")
	result.ExtraData = make(map[string]string)

	for _, selector := range selectors {
        result.ExtraData[selector] = doc.Find(selector).First().Text()
    }

	return result, nil
}

func pruneDocument(doc *goquery.Document, selectors []string) {
	for _, selector := range selectors {
		doc.Find(selector).Remove()
	}
}
