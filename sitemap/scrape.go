package sitemap

import (
	"fmt"
	"net/http"

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
	tokens := make(chan struct{}, concurrency) // To be use as Semaphore
	worklist := make(chan []string)
	dataResult := []Data{}

	go func() { worklist <- pages }()

	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if link != "" {
				n++
				go func(pageUrl string, tokens chan struct{}) {
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
					dataResult = append(dataResult, data)

				}(link, tokens)
			}
		}

	}
	return dataResult
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
