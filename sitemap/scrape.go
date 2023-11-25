package sitemap

import (
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/ngfenglong/go-scrape-flow/httpclient"
	"github.com/sirupsen/logrus"
)

// ------------ Data Structures ------------
type Data struct {
	URL             string
	Title           string
	MetaDescription string
	H1              string
	StatusCode      int
	ExtraData       map[string]string
}

type ContentMap struct {
	Tags    map[string]string
	Classes map[string]string
	IDs     map[string]string
}

// ------------ Global Variables ------------

var logger = log.New(logrus.InfoLevel, false)
var selectorsToRemove = []string{
	"body script",
	"body style",
	"body link",
}

// CrawlPages concurrently scrapes a list of pages and returns scraped data.
func CrawlPages(pages []string, concurrency int, selectors []string) []Data {
	var wg sync.WaitGroup
	tokens := make(chan struct{}, concurrency) // To be use as Semaphore
	dataChan := make(chan Data, len(pages))    // Using channel so dont have to use mutex

	for _, link := range pages {
		if link != "" {
			logger.Info("Scraping Page ", logrus.Fields{"Url": link})
			wg.Add(1)
			go scrapePage(link, tokens, dataChan, selectors, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(dataChan)
	}()

	return collectScrapedData(dataChan)
}

// scrapePage handles the scraping of a single page.
func scrapePage(pageUrl string, tokens chan struct{}, dataChan chan Data, selectors []string, wg *sync.WaitGroup) {
	defer wg.Done()
	tokens <- struct{}{}
	defer func() { <-tokens }()

	logger.Info("Requesting from URL ", logrus.Fields{"Url": pageUrl})
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
}

// collectScrapedData collects data from the data channel into a slice.
func collectScrapedData(dataChan chan Data) []Data {
	var finalResults []Data
	for data := range dataChan {
		finalResults = append(finalResults, data)
	}
	return finalResults
}

// ------------ Content Extraction Functions ------------

// extractContentFromResponse extracts content from an HTTP response.
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

func AnalysePageStructure(url string) ContentMap {
	res, err := httpclient.GetRequest(url)
	if err != nil {
		logger.Error("Error crawling page", logrus.Fields{"Url": url, "Error": err})
		return ContentMap{}
	}

	content, err := extractContentMapFromResponse(res)
	if err != nil {
		logger.Error("Error extracting content map from response", logrus.Fields{"Url": url, "Error": err})
	}

	return content
}

func extractContentMapFromResponse(res *http.Response) (ContentMap, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return ContentMap{}, err
	}

	// To remove unnecessary selectors if printing of the HTML is required
	pruneDocument(doc, selectorsToRemove)

	tags := make(map[string]string)
	classes := make(map[string]string)
	ids := make(map[string]string)

	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		// Get the tag name
		tagName := goquery.NodeName(s)
		tags[tagName] = doc.Find(tagName).First().Text()

		// Get all classes for this element and add them to the map
		class, exists := s.Attr("class")
		if exists {
			for _, className := range strings.Fields(class) { // Fields splits by space
				classes[className] = doc.Find(className).First().Text()
			}
		}

		// Get the ID for this element and add it to the map
		id, exists := s.Attr("id")
		if exists {
			ids[id] = doc.Find(id).First().Text()
		}
	})

	return ContentMap{
		Tags:    tags,
		Classes: classes,
		IDs:     ids,
	}, nil
}

func pruneDocument(doc *goquery.Document, selectors []string) {
	for _, selector := range selectors {
		doc.Find(selector).Remove()
	}
}
