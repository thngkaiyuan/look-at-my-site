package checker

type ResponseChecker interface {
	Check(domain string, resultChannel chan CheckerResult)
}

type CheckerResult struct {
	Ok          bool   `json:"ok"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
	results := make([]CheckerResult, len(c.checkers), len(c.checkers))
	resultChannel := make(chan CheckerResult)
	for _, ckr := range c.checkers {
		go ckr.Check(domain, resultChannel)
	}
	for i := range c.checkers {
		results[i] = <-resultChannel
	}
	return results
}
