package main

// run "go get golang.org/x/net/html" before importing "golang.org/x/net/html"
import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io"
	"strconv"
	"strings"
	"golang.org/x/net/html"
)

/*
	Crawler
*/

// define the use case of the crawler
func usage() {
	fmt.Fprintf(os.Stderr, "usage: crawler http://example.com\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		usage()
		fmt.Println("Please specify seed domain")
		os.Exit(1)
	}

	queue := make(chan string)
	
	filteredQueue := make(chan string)

	seedDomain := getDomain(args[0])
	fmt.Println("seed domain: ", seedDomain)

	go func() { queue <- args[0] }()
	go filterQueue(queue, filteredQueue)

	// pull from the filtered queue, add to the unfiltered queue
	for uri := range filteredQueue {
		enqueue(uri, queue, seedDomain)
	}
}

func filterQueue(in chan string, out chan string) {
	var seen = make(map[string]bool)
	for val := range in {
		if !seen[val] {
			seen[val] = true
			out <- val
		}
	}
}

func enqueue(uri string, queue chan string, seedDomain string) {
	fmt.Println("fetching", uri)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := all(resp.Body)

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if (absolute != "") && underDomain(absolute, seedDomain) {
			go func() { queue <- absolute }()
		}
	}
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func getDomain(uri string) string {
	domain := uri
	
	// Remove protocol
	if strings.Contains(domain, "http://") {
		domain = strings.TrimLeft(domain, "http://")
	} else if strings.Contains(uri, "https://") {
		domain = strings.TrimLeft(domain,"https://")
	} else {
		domain = ""
	}

	// Remove directory
	if strings.Contains(domain, "/") {
		domain = strings.TrimRight(strings.SplitAfter(domain, "/")[0], "/")
	}

	// Return domain
    return domain
}

func underDomain(uri string, domain string) bool {
	return strings.Contains(uri, domain)
}

/*
	Parser
*/

// All takes a reader object (like the one returned from http.Get())
// It returns a slice of strings representing the "href" attributes from
// anchor links found in the provided html.
// It does not close the reader passed to it.
func all(httpBody io.Reader) []string {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}
