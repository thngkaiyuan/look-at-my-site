package checker

type HstsChecker struct{}

func (c HstsChecker) Check(domains []string, resultChannel chan CheckerResult) {
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
		Title:            "This scan checks if your domain uses *HTTP Strict Transport Security (HSTS)*, which is a web security policy mechanism that allows web servers to declare that user agents should only interact with it using secure HTTPS connections.\nWithout HSTS, web servers may be vulnerable to protocol downgrade and cookie hijacking attacks. Web administrators are strongly encouraged to adopt HSTS on their web servers to thwart these attacks.",
		OkDescription:    "Safe because HSTS is enforced",
		OkUrls:           okUrls,
		NotOkDescription: "Not safe because HSTS is not enforced",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
