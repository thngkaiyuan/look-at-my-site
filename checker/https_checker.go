package checker

type HttpsChecker struct{}

func (c HttpsChecker) Check(domain string, resultChannel chan CheckerResult) {
	// Checks if we can connect to the domain using HTTPS

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
		Title:            "This scan checks if your domain supports *HTTPS*, a protocol for secure communication over a computer network. It protects the communication between your web server and its clients by means of encryption and authentication.\nWeb servers which do not support HTTPS are at risk of man-in-the-middle (MITM) attacks which include eavesdropping and tampering of communication contents.",
		OkDescription:    "Safe against MITM attacks if HTTPS is used",
		OkUrls:           okUrls,
		NotOkDescription: "Not safe against MITM attacks because HTTPS is not supported",
		NotOkUrls:        notOkUrls,
	}
	resultChannel <- result
}
