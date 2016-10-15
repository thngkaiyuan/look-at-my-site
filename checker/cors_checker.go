package checker

type CorsChecker struct{}

func (c CorsChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Do whatever you wanna check, and put result into the channel.
	// This is a stub result, you need to change it.
	result := CheckerResult{
		Ok:          false,
		Url:         "http://" + domain,
		Title:       "Dangerous CORS header(s) set",
		Description: "Access-Control-Allow-Origin is set to *, this will allow JavaScript from any site to send XHR requests to your site.",
	}
	resultChannel <- result
}
