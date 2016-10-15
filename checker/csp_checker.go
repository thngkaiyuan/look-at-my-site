package checker

type CspChecker struct{}

func (c CspChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Do whatever you wanna check, and put result into the channel.
	// This is a stub result, you need to change it.
	result := CheckerResult{
		Ok:          true,
		Url:         "http://" + domain,
		Title:       "CSP header is set correctly",
		Description: "A proper CSP header is set to prevent XSS.",
	}
	resultChannel <- result
}
