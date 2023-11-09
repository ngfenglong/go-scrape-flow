package httpclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

type RotatingProxies struct {
	proxies []string
	index   int
}

// NewRotatingProxies creates a new instance of RotatingProxies
func NewRotatingProxies(proxies []string) *RotatingProxies {
	return &RotatingProxies{
		proxies: proxies,
		index:   rand.Intn(len(proxies)), // Set the starting index to a random index
	}
}

func (rp *RotatingProxies) GetNextProxy() string {
	proxy := rp.proxies[rp.index]
	rp.index = (rp.index + 1) % len(rp.proxies)
	return proxy
}

func ReadProxiesFromFile(filename string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileBytes), "\n")
	var proxies []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			proxies = append(proxies, strings.TrimSpace(line))
		}
	}
	return proxies, nil
}

func CreateHTTPClientWithProxy(proxyAddr string) (*http.Client, error) {
	socks5URL := fmt.Sprintf("socks5://%s", proxyAddr)

	// Create a dialer from the proxy's address
	dialer, err := proxy.FromURL(parseURL(socks5URL), proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("failed to create dialer: %v", err)
	}

	// Setup HTTP transport and client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	return httpClient, nil
}

func parseURL(socks5ProxyAddr string) *url.URL {
	u, err := url.Parse(socks5ProxyAddr)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse proxy URL: %v", err))
	}
	return u
}

// region For testing Proxy
func TestProxy(proxyFile string) {
	proxies, err := ReadProxiesFromFile(proxyFile)
	if err != nil {
		log.Fatalf("Failed to read proxies: %v", err)
	}

	rp := NewRotatingProxies(proxies)

	for i := 0; i < len(proxies); i++ { // Rotate through all proxies once
		currentProxy := rp.GetNextProxy()
		logger.Info("Testing proxy", logrus.Fields{
			"index": i, "IP Address": currentProxy,
		})
		client, err := CreateHTTPClientWithProxy(currentProxy)
		if err != nil {
			logger.Error("Error creating client with proxy", logrus.Fields{"error": err})
			continue
		}
		ip, err := CheckPublicIP(client)
		if err != nil {
			logger.Error("Failed to fetch get request with proxy", logrus.Fields{"IP Address": currentProxy, "error": err})
			continue
		}

		if strings.HasPrefix(currentProxy, ip) {
			logger.Info("Proxy is valid", logrus.Fields{"IP Address": currentProxy})
		} else {
			logger.Error("Proxy is not working as expdcted", logrus.Fields{"IP Address": currentProxy})
		}
	}
}

func CheckPublicIP(httpClient *http.Client) (string, error) {
	resp, err := httpClient.Get("http://ipinfo.io/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}

//endregion
