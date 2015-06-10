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

	firstProcessor = chuper.ProcessorFunc(func(ctx *chuper.Context, doc *goquery.Document) bool {
		fmt.Printf("seed - %s - info: first %s %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL())
		return true
	})

	secondProcessor = chuper.ProcessorFunc(func(ctx *chuper.Context, doc *goquery.Document) bool {
		fmt.Printf("seed - %s - info: second %s %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL())
		return false

	})

	thirdProcessor = chuper.ProcessorFunc(func(ctx *chuper.Context, doc *goquery.Document) bool {
		fmt.Printf("seed - %s - info: third %s %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL())
		return true
	})
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = delay
	// crawler.CrawlPoliteness = true
	// crawler.Cache = nil
	// crawler.HTTPClient = prepareTorHTTPClient()

	crawler.Register(criteria, firstProcessor, secondProcessor, thirdProcessor)
	crawler.Start()

	crawler.Enqueue("GET", seeds...)
	crawler.Block()
}
