package httpclient

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	timeoutDuration = 10 * time.Second
	requestLimit    = 10
)

var userAgents = []string{
	// Chrome
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	// Firefox
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14.0; rv:109.0) Gecko/20100101 Firefox/118.0",
	"Mozilla/5.0 (X11; Linux i686; rv:109.0) Gecko/20100101 Firefox/118.0",
	// Safari
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
	// Internet Explorer 11
	"Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
}

var requestCounter int
var counterMutex sync.Mutex
var client = &http.Client{
	Timeout: timeoutDuration,
}

func incrementCounter() {
	counterMutex.Lock()
	requestCounter++
	counterMutex.Unlock()
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomDelay() {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
}

func randomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func GetRequest(url string) (*http.Response, error) {
	if requestCounter%requestLimit == 0 {
		client = &http.Client{Timeout: timeoutDuration}
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	randomDelay()
	incrementCounter()
	return client.Do(req)
}
