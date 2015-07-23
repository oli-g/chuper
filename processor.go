package chuper

import (
	"github.com/PuerkitoBio/goquery"
)

type Processor interface {
	Process(Context, *goquery.Document) bool
}

type ProcessorFunc func(Context, *goquery.Document) bool

func (p ProcessorFunc) Process(ctx Context, doc *goquery.Document) bool {
	return p(ctx, doc)
}
