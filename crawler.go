package chuper

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/oli-g/fetchbot"
)

// The Cmd struct embeds the basic fetchbot.Cmd implementation.
type Cmd struct {
	*fetchbot.Cmd
	S *url.URL
}

// Source returns the source of this command.
func (c *Cmd) Source() *url.URL {
	return c.S
}

type Crawler struct {
	f *fetchbot.Fetcher
	q *fetchbot.Queue
}

func defaultHandler() fetchbot.Handler {
	mux := fetchbot.NewMux()
	mux.Response().Method("GET").ContentType("text/html").Handler(fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		if err == nil {
			fmt.Printf("chuper - %s - info: [%d] %s %s - %s\n", time.Now().Format(time.RFC3339), res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL(), res.Header.Get("Content-Type"))
		}
	}))
	return mux
}

// New returns an initialized Crawler.
func New(d time.Duration, p bool) *Crawler {
	fetcher := fetchbot.New(defaultHandler())
	fetcher.CrawlDelay = d
	fetcher.CrawlPoliteness = p
	// fetcher.HttpClient = prepareTorHttpClient()

	return &Crawler{
		f: fetcher,
	}
}

func (c *Crawler) Start() *fetchbot.Queue {
	c.q = c.f.Start()
	return c.q
}

func (c *Crawler) Block() {
	c.q.Block()
}
