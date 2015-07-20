package chuper

import (
	"net/url"

	"github.com/PuerkitoBio/fetchbot"
)

type Context struct {
	*fetchbot.Context
	C Cache
}

func (c *Context) Cache() Cache {
	return c.C
}

func (c *Context) Queue() *fetchbot.Queue {
	return c.Q
}

func (c *Context) URL() *url.URL {
	return c.Cmd.URL()
}

func (c *Context) SourceURL() *url.URL {
	switch cmd := c.Cmd.(type) {
	case *Cmd:
		return cmd.SourceURL()
	case *CmdBasicAuth:
		return cmd.SourceURL()
	default:
		return nil
	}
}
