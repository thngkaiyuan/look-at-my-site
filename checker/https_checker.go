package checker

import (
	"errors"
)

type HttpsChecker struct{}

const (
	httpsTitle            = "This scan checks if your domain supports *HTTPS*, a protocol for secure communication over a computer network. It protects the confidentiality, integrity and authenticity of communications between your web server and its clients.\nWeb servers which do not support HTTPS are at risk of *man-in-the-middle (MITM) attacks* which include *eavesdropping* and *tampering of communication contents*."
	httpsOkDescription    = "Safe against MITM attacks if HTTPS is used"
	httpsNotOkDescription = "Not safe against MITM attacks because HTTPS is not supported"
)

func (c HttpsChecker) Check(in chan string, out chan CheckerResult) {
	okUrls := make([]string, 0)
	notOkUrls := make([]string, 0)
	okCh := make(chan string, 16)
	notOkCh := make(chan string, 16)

	count := 0
	for domain := range in {
		count++
		go checkHttps(domain, okCh, notOkCh)
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
		Title:            httpsTitle,
		OkDescription:    httpsOkDescription,
		OkUrls:           okUrls,
		NotOkDescription: httpsNotOkDescription,
		NotOkUrls:        notOkUrls,
	}

	if len(okUrls)+len(notOkUrls) == 0 {
		result.err = errors.New("All domains are down.")
	}

	out <- result
}

func checkHttps(domain string, okCh chan string, notOkCh chan string) {
	_, httpErr := httpClient.Head("http://" + domain)
	_, httpsErr := httpClient.Head("https://" + domain)

	if httpErr != nil && httpsErr != nil {
		// Domain is down, ignore.
		return
	}

	if httpsErr != nil {
		notOkCh <- domain
	} else {
		okCh <- domain
	}
}
