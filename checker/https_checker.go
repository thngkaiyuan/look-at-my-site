package checker

import (
	"errors"
	"net/http"
)

type HttpsChecker struct{}

const (
	title            = "HTTPS"
	okDescription    = "HTTPS is enabled for these domains."
	notOkDescription = "HTTPS is not enabled for these domains. Consider enabling HTTPS for security."
)

func (c HttpsChecker) Check(domains []string, resultChannel chan CheckerResult) {
	okUrls := make([]string, 0, len(domains))
	notOkUrls := make([]string, 0, len(domains))

	for _, domain := range domains {
		_, httpErr := http.Head("http://" + domain)
		_, httpsErr := http.Head("https://" + domain)

		if httpErr != nil && httpsErr != nil {
			// Domain is down, ignore.
			continue
		}

		if httpsErr != nil {
			notOkUrls = append(notOkUrls, domain)
		} else {
			okUrls = append(okUrls, domain)
		}
	}

	result := CheckerResult{
		Title:            title,
		OkDescription:    okDescription,
		OkUrls:           okUrls,
		NotOkDescription: notOkDescription,
		NotOkUrls:        notOkUrls,
	}

	if len(okUrls)+len(notOkUrls) == 0 {
		result.err = errors.New("All domains are down.")
	}

	resultChannel <- result
}
