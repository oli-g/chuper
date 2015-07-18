package chuper

import (
	"net/url"

	"github.com/PuerkitoBio/fetchbot"
)

type Command interface {
	URL() *url.URL
	Method() string
	SourceURL() *url.URL
}

type Cmd struct {
	*fetchbot.Cmd
	S *url.URL
}

func (c *Cmd) SourceURL() *url.URL {
	return c.S
}

type CmdBasicAuth struct {
	*fetchbot.Cmd
	S          *url.URL
	user, pass string
}

func (c *CmdBasicAuth) SourceURL() *url.URL {
	return c.S
}

func (c *CmdBasicAuth) BasicAuth() (user string, pwd string) {
	return c.user, c.pass
}
