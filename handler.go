package chuper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/oli-g/fetchbot"
)

var (
	DefaultErrorHandler = fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		fmt.Printf("chuper - %s - error: %s %s - %s\n", time.Now().Format(time.RFC3339), ctx.Cmd.Method(), ctx.Cmd.URL(), err)
	})

	DefaultLogHandlerFunc = func(ctx *fetchbot.Context, res *http.Response, err error) {
		if err == nil {
			fmt.Printf("chuper - %s - info: [%d] %s %s - %s\n", time.Now().Format(time.RFC3339), res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL(), res.Header.Get("Content-Type"))
		}
	}
)

func NewLogHandler(wrapped fetchbot.Handler, f func(ctx *fetchbot.Context, res *http.Response, err error)) fetchbot.Handler {
	return fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		f(ctx, res, err)
		wrapped.Handle(ctx, res, err)
	})
}
