package chuper

import (
	"net/url"

	"github.com/PuerkitoBio/fetchbot"
)

type Context interface {
	Cache() Cache
	Queue() Enqueuer
	URL() *url.URL
	SourceURL() *url.URL
}

type Ctx struct {
	*fetchbot.Context
	C Cache
}

func (c *Ctx) Cache() Cache {
	return c.C
}

func (c *Ctx) Queue() Enqueuer {
	return &Queue{c.Q}
}

func (c *Ctx) URL() *url.URL {
	return c.Cmd.URL()
}

func (c *Ctx) SourceURL() *url.URL {
	switch cmd := c.Cmd.(type) {
	case *Cmd:
		return cmd.SourceURL()
	case *CmdBasicAuth:
		return cmd.SourceURL()
	default:
		return nil
	}
}
