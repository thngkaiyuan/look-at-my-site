package crawler

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

type GoCrawlExtender struct {
	*gocrawl.DefaultExtender
	BaseDomain    string
	BufferChannel chan string
	Seen          map[string]bool
}

func (e *GoCrawlExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	host := ctx.URL().Host
	if _, ok := e.Seen[host]; !ok {
		e.Seen[host] = true
		select {
		case e.BufferChannel <- host:
		default:
		}
	}
	return nil, true
}

func (e *GoCrawlExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	host := ctx.URL().Host
	if isVisited {
		return false
	}
	if ctx.SourceURL() == nil || host == e.BaseDomain {
		return true
	}
	return strings.Contains(host, "."+e.BaseDomain) || strings.Contains(e.BaseDomain, host)
}

func (e *GoCrawlExtender) End(err error) {
	close(e.BufferChannel)
}
