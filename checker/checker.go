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
	checkers []ResponseChecker
}

func New() Checker {
	checkers := []ResponseChecker{
		new(HstsChecker),
		new(DnsRebindingChecker),
		new(CspChecker),
		new(CorsChecker),
	}
	return Checker{checkers}
}

func (c Checker) CheckAll(domain string) []CheckerResult {
	results := make([]CheckerResult, 0, len(c.checkers))
	resultChannel := make(chan CheckerResult)
	for _, ckr := range c.checkers {
		go ckr.Check(domain, resultChannel)
	}
	for range c.checkers {
		result := <-resultChannel
		if result.err != nil {
			continue
		}
		results = append(results, result)
	}
	return results
}
