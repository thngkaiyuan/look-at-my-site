package checker

type ResponseChecker interface {
	Check(domain string, resultChannel chan CheckerResult)
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
		new(HstsChecker),
		new(DnsRebindingChecker),
		new(CspChecker),
	}
	extraCheckers := []ResponseChecker{}
	return Checker{basicCheckers, extraCheckers}
}

func (c Checker) CheckBasic(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.basicCheckers))
	resultChannel := make(chan CheckerResult)
	for _, ckr := range c.basicCheckers {
		go ckr.Check(domain, resultChannel)
	}
	for range c.basicCheckers {
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
	resultChannel := make(chan CheckerResult)
	for _, ckr := range c.basicCheckers {
		go ckr.Check(domain, resultChannel)
	}
	for range c.basicCheckers {
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}
