package main

import (
	"fmt"
	"time"

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
)

func main() {
	crawler := chuper.New()
	crawler.CrawlDelay = delay
	// crawler.CrawlPoliteness = true
	// crawler.HTTPClient = prepareTorHttpClient()

	// crawler.Response(...).Register(handler1, handler2)

	crawler.Start()

	if err := crawler.Enqueue("GET", seeds...); err != nil {
		fmt.Printf("seed - %s - error: %s\n", time.Now().Format(time.RFC3339), err)
	}

	crawler.Block()
}
