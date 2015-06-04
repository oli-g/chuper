package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oli-g/chuper"
)

var (
	delay = 2 * time.Second

	seeds = []string{
		"http://www.gazzetta.it",
		"http://www.repubblica.it",
		"http://www.gazzetta.it",
		"http://www.repubblica.it",
		"http://www.corriere.it",
	}

	processor = chuper.ProcessorFunc(func(ctx *chuper.Context, doc *goquery.Document) error {
		fmt.Printf("seed - %s - process\n", time.Now().Format(time.RFC3339))
		return nil
	})
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = delay
	// crawler.CrawlPoliteness = true
	// crawler.Cache = nil
	// crawler.HTTPClient = prepareTorHTTPClient()

	crawler.Register(processor)

	crawler.Start()

	if err := crawler.Enqueue("GET", seeds...); err != nil {
		fmt.Printf("seed - %s - error: %s\n", time.Now().Format(time.RFC3339), err)
	}

	crawler.Block()
}
