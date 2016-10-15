package checker

type CorsChecker struct{}

func (c CorsChecker) Check(domain string, resultChannel chan CheckerResult) {
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
		Title:            "Cross-Origin Resource Sharing",
		OkDescription:    "The following subdomains have set a proper CORS header.",
		OkUrls:           okUrls,
		NotOkDescription: "The following subdomains have their CORS header set to *. They might be vulnerable to *Cross-Site Request Forgery(CSRF)*.",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
