package checker

type CspChecker struct{}

func (c CspChecker) Check(domain string, resultChannel chan CheckerResult) {
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
		Title:            "Content-Secuity-Policy",
		OkDescription:    "The following subdomains have their CSP header properly set.",
		OkUrls:           okUrls,
		NotOkDescription: "The following subdomains do not set CSP header or have their CSP header set to \"inline\". They might be vulnerable to *Cross-Site-Scripting(XSS)* attacks.",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
