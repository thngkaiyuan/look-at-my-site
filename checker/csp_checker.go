package checker

type CspChecker struct{}

func (c CspChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Check for presence of CSP and absence of 'unsafe-inline' in the CSP header
	
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
		Title:            "This scan checks the *Content-Secuity-Policy (CSP)* of your domain(s). The CSP response header helps you reduce cross-site scripting (XSS) risks on modern browsers by declaring what dynamic resources are allowed to load via the response header.\nWeb servers which do not use CSP or which support inline scripts are at risk of cross-site scripting (XSS), clickjacking and other code injection attacks.",
		OkDescription:    "Safe against most XSS attacks",
		OkUrls:           okUrls,
		NotOkDescription: "Not safe because they either lack CSP or support 'unsafe-inline' scripts",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
