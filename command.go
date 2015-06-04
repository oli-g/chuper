package chuper

import (
	"net/url"

	"github.com/oli-g/fetchbot"
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

func NewCmd(u *url.URL, m string, s *url.URL) *Cmd {
	return &Cmd{
		&fetchbot.Cmd{U: u, M: m},
		s,
	}
}
