package chuper

import (
	"net/url"

	"github.com/oli-g/fetchbot"
)

type Context struct {
	*fetchbot.Context
	C Cache
}

func (c *Context) SourceURL() *url.URL {
	switch cmd := c.Cmd.(type) {
	case Cmd:
		return cmd.SourceURL()
	}
	return nil
}
