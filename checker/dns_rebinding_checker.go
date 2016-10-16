package checker

type DnsRebindingChecker struct{}

func (c DnsRebindingChecker) Check(domains []string, resultChannel chan CheckerResult) {
	// If the domain does not support HTTP (i.e. only supports HTTPS), then it is secure
	// Otherwise we check if we are able to connect to the domain using an invalid Host field in the header

	// Do whatever you wanna check, and put result into the channel.
	// This is a stub result, you need to change it.
	okUrls := []string{
		"google.com",
		"nus.edu.sg",
	}
	notOkUrls := []string{
		"twitter.com",
		"youtube.com",
	}
	result := CheckerResult{
		Title:            "In this scan, our server will attempt to connect to your domain(s) using an invalid *Host header*. If it is successful in doing so, then your domain might be used as a target of DNS rebinding attacks.\nWeb servers which reject HTTP requests with unrecognized Host headers or which strictly requires HTTPS connections are safe against DNS rebinding attacks.",
		OkDescription:    "Safe either because the web server rejects unrecognized Host headers or strictly requires HTTPS",
		OkUrls:           okUrls,
		NotOkDescription: "Unsafe because the invalid Host header was ignored and HTTP connections are supported",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
