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

	criteria = &chuper.ResponseCriteria{
		Method:      "GET",
		ContentType: "text/html",
		Status:      200,
		Host:        "www.gazzetta.it",
	}

	processor = chuper.ProcessorFunc(func(ctx *chuper.Context, doc *goquery.Document) error {
		fmt.Printf("seed - %s - info: processing %s %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL())
		return nil
	})
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = delay
	// crawler.CrawlPoliteness = true
	// crawler.Cache = nil
	// crawler.HTTPClient = prepareTorHTTPClient()

	crawler.Register(criteria, processor)
	crawler.Start()

	crawler.Enqueue("GET", seeds...)
	crawler.Block()
}
