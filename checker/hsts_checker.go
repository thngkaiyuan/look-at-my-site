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
		Title:            "This scan checks if your domain enforces HTTPS connections using *HTTP Strict Transport Security (HSTS)*. Without enforcing HSTS, an attacker with man-in-the-middle capability could potentially perform HTTPS downgrade attacks to compromise communications between your web server and its clients.\nVulnerable web servers should...",
		OkDescription:    "Safe because HSTS is enforced",
		OkUrls:           okUrls,
		NotOkDescription: "Vulnerable because HSTS is not enforced",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
