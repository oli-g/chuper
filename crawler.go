package chuper

import (
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
)

const (
	DefaultCrawlDelay      = 5 * time.Second
	DefaultCrawlPoliteness = false
	DefaultLogFormat       = "text"
	DefaultLogLevel        = "info"
	DefaultUserAgent       = fetchbot.DefaultUserAgent
)

var (
	DefaultHTTPClient = http.DefaultClient
	DefaultCache      = NewMemoryCache()
)

type Crawler struct {
	CrawlDelay      time.Duration
	CrawlDuration   time.Duration
	CrawlPoliteness bool
	LogFormat       string
	LogLevel        string
	Logger          *logrus.Logger
	UserAgent       string
	HTTPClient      fetchbot.Doer
	Cache           Cache

	mux *fetchbot.Mux
	f   *fetchbot.Fetcher
	q   *fetchbot.Queue
}

// New returns an initialized Crawler.
func New() *Crawler {
	return &Crawler{
		CrawlDelay:      DefaultCrawlDelay,
		CrawlPoliteness: DefaultCrawlPoliteness,
		LogFormat:       DefaultLogFormat,
		LogLevel:        DefaultLogLevel,
		UserAgent:       DefaultUserAgent,
		HTTPClient:      DefaultHTTPClient,
		Cache:           DefaultCache,
		mux:             fetchbot.NewMux(),
	}
}

func (c *Crawler) Start() Enqueuer {
	if c.Logger == nil {
		c.Logger = newLogger(c.LogFormat, c.LogLevel)
	}

	c.mux.HandleErrors(c.newErrorHandler())
	h := c.newRequestHandler()

	f := fetchbot.New(h)
	f.CrawlDelay = c.CrawlDelay
	f.DisablePoliteness = !c.CrawlPoliteness
	f.HttpClient = c.HTTPClient
	f.UserAgent = c.UserAgent

	c.f = f
	c.q = c.f.Start()

	if c.CrawlDuration > 0 {
		go func() {
			t := time.After(c.CrawlDuration)
			<-t
			c.q.Close()
		}()
	}

	return &Queue{c.q}
}

func (c *Crawler) Block() {
	c.q.Block()
}

func (c *Crawler) Finish() {
	c.q.Close()
}

type ResponseCriteria struct {
	Method      string
	ContentType string
	Status      int
	MinStatus   int
	MaxStatus   int
	Path        string
	Host        string
}

func (c *Crawler) Match(r *ResponseCriteria) *fetchbot.ResponseMatcher {
	m := c.mux.Response()

	if r.Method != "" {
		m.Method(r.Method)
	}

	if r.ContentType != "" {
		m.ContentType(r.ContentType)
	}

	if r.Status != 0 {
		m.Status(r.Status)
	} else {
		if r.MinStatus != 0 && r.MaxStatus != 0 {
			m.StatusRange(r.MinStatus, r.MaxStatus)
		} else {
			if r.MinStatus != 0 {
				m.Status(r.MinStatus)
			}
			if r.MaxStatus != 0 {
				m.Status(r.MaxStatus)
			}
		}
	}

	if r.Path != "" {
		m.Path(r.Path)
	}

	if r.Host != "" {
		m.Host(r.Host)
	}

	return m
}

func (c *Crawler) Register(rc *ResponseCriteria, procs ...Processor) {
	m := c.Match(rc)
	h := c.newHTMLHandler(procs...)
	m.Handler(h)
}

func newLogger(format, level string) *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = newFormatter(format)
	log.Level = parseLogLevel(level)
	return log
}

func newFormatter(format string) logrus.Formatter {
	switch format {
	case "text", "":
		return &logrus.TextFormatter{}
	case "json":
		return &logrus.JSONFormatter{}
	default:
		return &logrus.TextFormatter{}
	}
}

func parseLogLevel(level string) logrus.Level {
	switch level {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}

func (c *Crawler) newErrorHandler() fetchbot.Handler {
	return fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		c.Logger.WithFields(logrus.Fields{
			"url":    ctx.Cmd.URL(),
			"method": ctx.Cmd.Method(),
		}).Error(err)
	})
}

func (c *Crawler) newRequestHandler() fetchbot.Handler {
	return fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		if res != nil {
			context := &Ctx{ctx, c.Cache, c.Logger}
			c.Logger.WithFields(logrus.Fields{
				"method":       context.Method(),
				"status":       res.StatusCode,
				"content_type": res.Header.Get("Content-Type"),
				"depth":        context.Depth(),
			}).Info(context.URL())
		}
		c.mux.Handle(ctx, res, err)
	})
}

func (c *Crawler) newHTMLHandler(procs ...Processor) fetchbot.Handler {
	return fetchbot.HandlerFunc(func(ctx *fetchbot.Context, res *http.Response, err error) {
		context := &Ctx{ctx, c.Cache, c.Logger}
		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			c.Logger.WithFields(logrus.Fields{
				"url":    context.URL(),
				"method": context.Method(),
			}).Error(err)
			return
		}

		for _, p := range procs {
			ok := p.Process(context, doc)
			if !ok {
				return
			}
		}
	})
}
