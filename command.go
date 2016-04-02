package chuper

import (
	"net/url"

	"github.com/PuerkitoBio/fetchbot"
)

type Command interface {
	URL() *url.URL
	Method() string
	SourceURL() *url.URL
	Depth() int
}

type Cmd struct {
	*fetchbot.Cmd
	S *url.URL
	D int
}

func (c *Cmd) SourceURL() *url.URL {
	return c.S
}

func (c *Cmd) Depth() int {
	return c.D
}

type CmdBasicAuth struct {
	*fetchbot.Cmd
	S          *url.URL
	D          int
	user, pass string
}

func (c *CmdBasicAuth) SourceURL() *url.URL {
	return c.S
}

func (c *CmdBasicAuth) Depth() int {
	return c.D
}

func (c *CmdBasicAuth) BasicAuth() (string, string) {
	return c.user, c.pass
}
