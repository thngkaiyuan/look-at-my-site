package checker

import (
	"log"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/purell"
	"github.com/thngkaiyuan/look-at-my-site/crawler"
)

type ResponseChecker interface {
	Check(in chan string, out chan CheckerResult)
}

type CheckerResult struct {
	Title            string   `json:"title"`
	OkDescription    string   `json:"okDescription"`
	NotOkDescription string   `json:"notOkDescription"`
	OkUrls           []string `json:"okUrls"`
	NotOkUrls        []string `json:"notOkUrls"`
	err              error
}

type Checker struct {
	basicCheckers []ResponseChecker
	extraCheckers []ResponseChecker
}

func New() Checker {
	basicCheckers := []ResponseChecker{
		new(HttpsChecker),
	}
	extraCheckers := []ResponseChecker{}
	return Checker{basicCheckers, extraCheckers}
}

func (c Checker) CheckBasic(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.basicCheckers))

	domainChannels := make([]chan string, len(c.basicCheckers))
	resultChannel := make(chan CheckerResult)

	for i, ckr := range c.basicCheckers {
		domainChannels[i] = make(chan string)
		go ckr.Check(domainChannels[i], resultChannel)
	}

	go func() {
		time.Sleep(5 * time.Second)
		for _, ch := range domainChannels {
			close(ch)
		}
	}()

	for i := range c.basicCheckers {
		domainChannels[i] <- domain
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}

func (c Checker) CheckAll(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.basicCheckers))

	bufferChannel := make(chan string, 16)
	domainChannels := make([]chan string, len(c.basicCheckers))
	resultChannel := make(chan CheckerResult)

	for i, ckr := range c.basicCheckers {
		domainChannels[i] = make(chan string, 16)
		go ckr.Check(domainChannels[i], resultChannel)
	}

	ext := &crawler.GoCrawlExtender{
		&gocrawl.DefaultExtender{},
		domain,
		bufferChannel,
		make(map[string]bool),
	}

	opts := gocrawl.NewOptions(ext)

	opts.CrawlDelay = 1 * time.Second
	opts.SameHostOnly = false
	opts.URLNormalizationFlags = purell.FlagsUsuallySafeGreedy
	opts.LogFlags = gocrawl.LogNone

	crawler := gocrawl.NewCrawlerWithOptions(opts)

	seeds := []string{
		"http://" + domain,
		"https://" + domain,
	}

	log.Println("Starting crawler...")
	go crawler.Run(seeds)

	timeoutCh := time.After(30 * time.Second)
	stop := false
	for !stop {
		select {
		case domain := <-bufferChannel:
			for _, ch := range domainChannels {
				ch <- domain
			}
		case <-timeoutCh:
			log.Println("Time's up! Finalizing...")
			for _, ch := range domainChannels {
				close(ch)
			}
			stop = true
			go crawler.Stop()
		}
	}

	for range c.basicCheckers {
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}
