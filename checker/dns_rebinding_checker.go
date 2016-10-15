package checker

type DnsRebindingChecker struct{}

func (c DnsRebindingChecker) Check(domain string, resultChannel chan CheckerResult) {
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
		Title:            "Validation of \"Host\" field",
		OkDescription:    "The following subdomains checks the \"Host\" field in the HTTP(S) request header.",
		OkUrls:           okUrls,
		NotOkDescription: "The following subdomains ignores the \"Host\" field of the HTTP(S) request header. They might be vulnerable to *DNS Rebinding Attack*.",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
