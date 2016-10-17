package checker

type ResponseChecker interface {
	Check(in chan string, out chan CheckerResult)
}

type CheckerResult struct {
	Title            string   `json:"title"`
	OkDescription    string   `json:"okDescription"`
	NotOkDescription string   `json:"notOkDescription"`
	OkUrls           []string `json:"okUrls"`
	NotOkUrls        []string `json:"notOkUrls"`
	err              error
}

type Checker struct {
	basicCheckers []ResponseChecker
	extraCheckers []ResponseChecker
}

func New() Checker {
	basicCheckers := []ResponseChecker{
		new(HttpsChecker),
	}
	extraCheckers := []ResponseChecker{}
	return Checker{basicCheckers, extraCheckers}
}

func (c Checker) CheckBasic(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.basicCheckers))

	domainChannels := make([]chan string, len(c.basicCheckers))
	resultChannel := make(chan CheckerResult)

	for i, ckr := range c.basicCheckers {
		domainChannels[i] = make(chan string)
		go ckr.Check(domainChannels[i], resultChannel)
	}

	for i := range c.basicCheckers {
		domainChannels[i] <- domain
		close(domainChannels[i])
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}

func (c Checker) CheckAll(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.basicCheckers))

	domainChannels := make([]chan string, len(c.basicCheckers))
	resultChannel := make(chan CheckerResult)

	for i, ckr := range c.basicCheckers {
		domainChannels[i] = make(chan string)
		go ckr.Check(domainChannels[i], resultChannel)
	}
	for i := range c.basicCheckers {
		domainChannels[i] <- domain
		close(domainChannels[i])
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}
