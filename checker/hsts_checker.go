package checker

import (
	"errors"
	"net/http"
)

type HstsChecker struct{}

const (
	hstsTitle            = "This scan checks if your domain uses *HTTP Strict Transport Security (HSTS)*, which is a web security policy mechanism that allows web servers to declare that user agents should only interact with it using secure HTTPS connections.\nWithout HSTS, web servers may be vulnerable to *protocol downgrade* and *cookie hijacking attacks*. Web administrators are strongly encouraged to adopt HSTS on their web servers to thwart these attacks."
	hstsOkDescription    = "Safe because HSTS is enforced"
	hstsNotOkDescription = "Not safe because HSTS is not enforced"
)

func (c HstsChecker) Check(in chan string, out chan CheckerResult) {
	okUrls := make([]string, 0)
	notOkUrls := make([]string, 0)
	okCh := make(chan string, 16)
	notOkCh := make(chan string, 16)

	count := 0
	for domain := range in {
		count++
		go checkHsts(domain, okCh, notOkCh)
	}

	for i := 0; i < count; i++ {
		select {
		case domain := <-okCh:
			okUrls = append(okUrls, domain)
		case domain := <-notOkCh:
			notOkUrls = append(notOkUrls, domain)
		default:
		}
	}

	result := CheckerResult{
		Title:            hstsTitle,
		OkDescription:    hstsOkDescription,
		OkUrls:           okUrls,
		NotOkDescription: hstsNotOkDescription,
		NotOkUrls:        notOkUrls,
	}

	if len(okUrls)+len(notOkUrls) == 0 {
		result.err = errors.New("All domains are down.")
	}

	out <- result
}

func checkHsts(domain string, okCh chan string, notOkCh chan string) {
	resp, httpsErr := httpClient.Head("https://" + domain)

	if httpsErr != nil {
		// Domain is down, ignore.
		return
	}

	hstsHeader := resp.Header.Get(http.CanonicalHeaderKey("Strict-Transport-Security"))

	if hstsHeader == "" {
		notOkCh <- domain
	} else {
		okCh <- domain
	}
}
