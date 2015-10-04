package main

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oli-g/chuper"
)

var (
	delay = 2 * time.Second

	depth = 0

	seeds = []string{
		"http://www.repubblica.it",
		"http://www.corriere.it",
		"http://www.repubblica.it",
		"http://www.corriere.it",
		"http://www.gazzetta.it",
	}

	criteria = &chuper.ResponseCriteria{
		Method:      "GET",
		ContentType: "text/html",
		Status:      200,
		Host:        "www.gazzetta.it",
	}

	firstProcessor = chuper.ProcessorFunc(func(ctx chuper.Context, doc *goquery.Document) bool {
		ctx.Log(map[string]interface{}{
			"url":    ctx.URL().String(),
			"source": ctx.SourceURL().String(),
		}).Info("First processor")
		return true
	})

	secondProcessor = chuper.ProcessorFunc(func(ctx chuper.Context, doc *goquery.Document) bool {
		ctx.Log(map[string]interface{}{
			"url":    ctx.URL().String(),
			"source": ctx.SourceURL().String(),
		}).Info("Second processor")
		return false

	})

	thirdProcessor = chuper.ProcessorFunc(func(ctx chuper.Context, doc *goquery.Document) bool {
		ctx.Log(map[string]interface{}{
			"url":    ctx.URL().String(),
			"source": ctx.SourceURL().String(),
		}).Info("Third processor")
		return true
	})
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = delay

	crawler.Register(criteria, firstProcessor, secondProcessor, thirdProcessor)
	q := crawler.Start()

	for _, u := range seeds {
		q.Enqueue("GET", u, "www.google.com", depth)
		depth++
	}

	crawler.Finish()
}
