package chuper

import (
	"net/http"
	"time"

	"github.com/oli-g/fetchbot"
)

const (
	DefaultCrawlDelay      = 5 * time.Second
	DefaultCrawlPoliteness = false
)

var (
	DefaultHTTPClient = http.DefaultClient
)

type Crawler struct {
	CrawlDelay      time.Duration
	CrawlPoliteness bool
	HTTPClient      fetchbot.Doer
	Cache           Cache
	ErrorHandler    fetchbot.Handler
	LogHandlerFunc  func(ctx *fetchbot.Context, res *http.Response, err error)

	mux *fetchbot.Mux
	f   *fetchbot.Fetcher
	q   *fetchbot.Queue
}

// New returns an initialized Crawler.
func New() *Crawler {
	return &Crawler{
		CrawlDelay:      DefaultCrawlDelay,
		CrawlPoliteness: DefaultCrawlPoliteness,
		HTTPClient:      DefaultHTTPClient,
		Cache:           DefaultCache,
		ErrorHandler:    DefaultErrorHandler,
		LogHandlerFunc:  DefaultLogHandlerFunc,
		mux:             fetchbot.NewMux(),
	}
}

// TODO: l'Handler chiama una successione di subhandlers che ritornano true/false.
// Definire quindi un ProcessorFunc type

func (c *Crawler) Start() *fetchbot.Queue {
	c.mux.HandleErrors(c.ErrorHandler)
	l := NewLogHandler(c.mux, c.LogHandlerFunc)
	f := fetchbot.New(l)
	// h := crawlerHandler(c.Cache, c.ScraperHandler, c.EnqueuerHandler)

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

func (c *Crawler) Enqueue(method string, rawurl ...string) error {
	for _, u := range rawurl {
		ok := true
		if c.mustCache() {
			ok, _ = c.Cache.SetNX(u, true)
		}
		if ok {
			if _, err := c.q.SendString(method, u); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Crawler) mustCache() bool {
	if c.Cache == nil {
		return false
	}
	return true
}
