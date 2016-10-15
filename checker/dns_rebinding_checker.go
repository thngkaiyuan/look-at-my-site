package checker

type DnsRebindingChecker struct{}

func (c DnsRebindingChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Do whatever you wanna check, and put result into the channel.
	// This is a stub result, you need to change it.
	result := CheckerResult{
		Ok:          false,
		Url:         "http://" + domain,
		Title:       "Host field is being ignored.",
		Description: "The server is ignoring the \"Host\" field in the HTTP request header. DNS rebinding attack might be possible.",
	}
	resultChannel <- result
}
