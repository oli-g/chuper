package chuper

import "github.com/PuerkitoBio/goquery"

type Processor interface {
	Process(Context, []byte, *goquery.Document) bool
}

type ProcessorFunc func(Context, []byte, *goquery.Document) bool

func (p ProcessorFunc) Process(ctx Context, body []byte, doc *goquery.Document) bool {
	return p(ctx, body, doc)
}
