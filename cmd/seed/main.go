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
	crawler := chuper.New(crawlDelay, false)
	queue := crawler.Start()

	if _, err := queue.SendStringGet(seed); err != nil {
		fmt.Printf("chuper - error: %s\n", err)
	}

	crawler.Block()
}
