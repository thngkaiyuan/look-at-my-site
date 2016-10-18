package checker

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type DnsRebindingChecker struct{}

const (
	dnsRebindingTitle            = "In this scan, our server will attempt to connect to your domain(s) using an invalid *Host header*. If it is successful in doing so, then your domain might be used as a target of DNS rebinding attacks.\nWeb servers which reject HTTP requests with unrecognized Host headers or which strictly requires HTTPS connections are safe against DNS rebinding attacks."
	dnsRebindingOkDescription    = "Safe either because the web server rejects unrecognized Host headers or strictly requires HTTPS"
	dnsRebindingNotOkDescription = "Not safe because either the invalid Host header was ignored or HTTP connections are supported"
)

func (c DnsRebindingChecker) Check(in chan string, out chan CheckerResult) {
	okUrls := make([]string, 0)
	notOkUrls := make([]string, 0)
	okCh := make(chan string, 16)
	notOkCh := make(chan string, 16)

	count := 0
	for domain := range in {
		count++
		go checkDnsRebinding(domain, okCh, notOkCh)
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
		Title:            dnsRebindingTitle,
		OkDescription:    dnsRebindingOkDescription,
		OkUrls:           okUrls,
		NotOkDescription: dnsRebindingNotOkDescription,
		NotOkUrls:        notOkUrls,
	}

	if len(okUrls)+len(notOkUrls) == 0 {
		result.err = errors.New("All domains are down.")
	}

	out <- result
}

func checkDnsRebinding(domain string, okCh chan string, notOkCh chan string) {
	// If the domain does not support HTTP (i.e. only supports HTTPS), then it is secure
	// Otherwise we check if we are able to connect to the domain using an invalid Host field in the header

	_, err := httpClient.Head("http://" + domain)

	if err != nil {
		_, err = httpClient.Head("https://" + domain)

		if err != nil {
			// Domain is down, ignore.
			return
		} else {
			// Only support HTTPS, safe from DNS Rebinding Attack
			okCh <- domain
		}
	}

	// Try to connect to the domain using an invalid Host field in the header
	fakeReq, _ := http.NewRequest("GET", "http://" + domain, nil)
	fakeReq.Host = "look-at-my.site"
	fakeResp, err := httpClient.Do(fakeReq)
	if err != nil {
		return
	}

	fakeContent, err := ioutil.ReadAll(fakeResp.Body)
	defer fakeResp.Body.Close()
	if err != nil {
		return
	}

	// Try to connect to the domain without invalid Host field in the header
	realReq, _ := http.NewRequest("GET", "http://" + domain, nil)
	realResp, err := httpClient.Do(realReq)
	if err != nil {
		return 
	}

	realContent, err := ioutil.ReadAll(realResp.Body)
	defer realResp.Body.Close()
	if err != nil {
		return
	}
	
	// Compare the result of the 2 requests
	if (string(fakeContent) == string(realContent)) {
		// The results are similar, unsafe from DNS Rebinding Attack
		notOkCh <- domain
	} else {
		// The results are not similar, safe from DNS Rebinding Attack
		okCh <- domain
	}
}
