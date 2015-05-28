package chuper

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/oli-g/fetchbot"
)

const (
	DefaultCrawlPoliteness = false
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
	CrawlDelay      time.Duration
	CrawlPoliteness bool
	HTTPClient      fetchbot.Doer

	f *fetchbot.Fetcher
	q *fetchbot.Queue
}

func errorHandler(ctx *fetchbot.Context, res *http.Response, err error) {
	fmt.Printf("chuper - %s - error: %s %s - %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL(), err)
}

func Handler() fetchbot.Handler {
	mux := fetchbot.NewMux()
	mux.HandleErrors(fetchbot.HandlerFunc(errorHandler))

	mux.Response().Method("GET").ContentType("text/html").Handler(fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		if err == nil {
			fmt.Printf("chuper - %s - info: [%d] %s %s - %s\n", time.Now().Format(time.RFC3339), res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL(), res.Header.Get("Content-Type"))
		}
	}))

	return mux
}

// New returns an initialized Crawler.
func New(d time.Duration) *Crawler {
	return &Crawler{
		CrawlDelay:      d,
		CrawlPoliteness: DefaultCrawlPoliteness,
	}
}

func (c *Crawler) Start() *fetchbot.Queue {
	f := fetchbot.New(Handler())

	f.CrawlDelay = c.CrawlDelay
	f.CrawlPoliteness = c.CrawlPoliteness

	if c.HTTPClient != nil {
		f.HttpClient = c.HTTPClient
	}

	c.f = f
	c.q = c.f.Start()

	return c.q
}

func (c *Crawler) Block() {
	c.q.Block()
}
