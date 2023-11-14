package main

import (
	"os"

	"github.com/ngfenglong/go-scrape-flow/cmd"
	"github.com/ngfenglong/go-scrape-flow/helper/log"
	"github.com/sirupsen/logrus"
)

var logger = log.New(logrus.InfoLevel, false)

func main() {
	// Provided a sitemap URL
	if err := cmd.Execute(); err != nil {
		logger.Error("Error executing command", logrus.Fields{"error": err})
		os.Exit(1)
	}
}
