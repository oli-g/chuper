package chuper

import (
	"github.com/oli-g/fetchbot"
)

type Context struct {
	*fetchbot.Context
	C Cache
}
