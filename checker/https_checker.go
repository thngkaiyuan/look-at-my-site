package checker

import (
	"errors"
	"net/http"
)

type HttpsChecker struct{}

const (
	title            = "This scan checks if your domain supports *HTTPS*, a protocol for secure communication over a computer network. It protects the communication between your web server and its clients by means of encryption and authentication.\nWeb servers which do not support HTTPS are at risk of man-in-the-middle (MITM) attacks which include eavesdropping and tampering of communication contents."
	okDescription    = "Safe against MITM attacks if HTTPS is used"
	notOkDescription = "Not safe against MITM attacks because HTTPS is not supported"
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
