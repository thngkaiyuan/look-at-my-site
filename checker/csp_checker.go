package checker

import (
	"errors"
	"net/http"
	"strings"
)

type CspChecker struct{}

const (
	cspTitle            = "This scan checks the *Content-Secuity-Policy (CSP)* of your domain(s). The CSP response header helps you reduce cross-site scripting (XSS) risks on modern browsers by declaring what dynamic resources are allowed to load via the response header.\nWeb servers which do not use CSP or which support inline scripts are at risk of XSS, clickjacking and other code injection attacks."
	cspOkDescription    = "Safe against known XSS attacks"
	cspNotOkDescription = "Not safe because they either lack CSP or support 'unsafe-inline' scripts"
)

func (c CspChecker) Check(in chan string, out chan CheckerResult) {
	okUrls := make([]string, 0)
	notOkUrls := make([]string, 0)
	okCh := make(chan string, 16)
	notOkCh := make(chan string, 16)

	count := 0
	for domain := range in {
		count++
		go checkCsp(domain, okCh, notOkCh)
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
		Title:            cspTitle,
		OkDescription:    cspOkDescription,
		OkUrls:           okUrls,
		NotOkDescription: cspNotOkDescription,
		NotOkUrls:        notOkUrls,
	}

	if len(okUrls)+len(notOkUrls) == 0 {
		result.err = errors.New("All domains are down.")
	}

	out <- result
}

func checkCsp(domain string, okCh chan string, notOkCh chan string) {
	resp, err := httpClient.Head("http://" + domain)

	if err != nil {
		resp, err = httpClient.Head("https://" + domain)

		if err != nil {
			// Domain is down, ignore.
			return
		}
	}

	cspHeader := resp.Header.Get(http.CanonicalHeaderKey("Content-Security-Policy"))
	if cspHeader == "" || strings.Contains(cspHeader, "unsafe") {
		notOkCh <- domain
	} else {
		okCh <- domain
	}
}
