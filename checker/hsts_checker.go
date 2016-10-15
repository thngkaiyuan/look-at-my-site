package checker

type HstsChecker struct{}

func (c HstsChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Do whatever you wanna check, and put result into the channel.
	// This is a stub result, you need to change it.
	result := CheckerResult{
		Ok:          false,
		Url:         "http://" + domain,
		Title:       "HSTS header is not set.",
		Description: "Your website might be vulnerable to HTTPS downgrading attack or MITM attack.",
	}
	resultChannel <- result
}
