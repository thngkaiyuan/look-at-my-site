package checker

type HstsChecker struct{}

func (c HstsChecker) Check(domain string, resultChannel chan CheckerResult) {
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
		Title:            "HTTP Strict Transport Security (HSTS)",
		OkDescription:    "The following subdomains have set HSTS header.",
		OkUrls:           okUrls,
		NotOkDescription: "The following subdomains do not set HSTS header. They might be vulnerable to *HTTPS Downgrade Attack* and *Man-In-The-Middle(MITM) Attack*.",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
