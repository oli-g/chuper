package chuper

import (
	"github.com/PuerkitoBio/goquery"
)

type Processor interface {
	Process(*Context, *goquery.Document) error
}

type ProcessorFunc func(*Context, *goquery.Document) error

func (p ProcessorFunc) Process(ctx *Context, doc *goquery.Document) error {
	return p(ctx, doc)
}
