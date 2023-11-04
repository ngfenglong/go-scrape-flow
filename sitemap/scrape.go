package sitemap

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/ngfenglong/go-scrape-flow/httpclient"
)

type Data struct {
	URL             string
	Title           string
	MetaDescription string
	H1              string
	StatusCode      int
	// Add individual type
}

var selectorsToRemove = []string{
	"body script",
	"body style",
	"body link",
}

func CrawlPages(pages []string, concurrency int) []Data {
	var wg sync.WaitGroup
	tokens := make(chan struct{}, concurrency) // To be use as Semaphore
	dataChan := make(chan Data, len(pages))    // Using channel so dont have to use mutex

	for _, link := range pages {
		if link != "" {
			fmt.Println("Checking Url - ", link)
			wg.Add(1)
			go func(pageUrl string, tokens chan struct{}) {
				defer wg.Done()
				tokens <- struct{}{}
				fmt.Println("Requesting from URL: ", pageUrl) // to add into log
				<-tokens
				res, err := httpclient.GetRequest(pageUrl)
				if err != nil {
					fmt.Printf("Error: %v", err) // to add into log
					return
				}

				data, err := extractContentFromResponse(res)
				if err != nil {
					fmt.Printf("Error: %v", err) // to add into log
					return
				}
				fmt.Println(data)
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

func extractContentFromResponse(res *http.Response) (Data, error) {
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

	return result, nil
}

func pruneDocument(doc *goquery.Document, selectors []string) {
	for _, selector := range selectors {
		doc.Find(selector).Remove()
	}
}
