package main

import (
	"fmt"
	"time"

	"github.com/oli-g/chuper"
)

var (
	crawlDelay = 2 * time.Second
	seed       = "http://www.gazzetta.it"
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = crawlDelay
	// crawler.CrawlPoliteness = true
	// crawler.HTTPClient = prepareTorHttpClient()

	// crawler.Response(...).Register(handler1, handler2)

	queue := crawler.Start()

	if _, err := queue.SendStringGet(seed); err != nil {
		fmt.Printf("seed - %s - error: %s\n", time.Now().Format(time.RFC3339), err)
	}

	crawler.Block()
}
