package main

import (
	"fmt"
	"sync"
	"log"
	"os"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type ConcurrentMap struct {
	visited map[string]bool
	mux sync.Mutex
}

var cmap ConcurrentMap
var logger log.Logger
var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	// Base case: depth <= 0
	if depth <= 0 {
		return
	}
	// Base case: URL already visited
	cmap.mux.Lock()
	_, exists := cmap.visited[url]
	cmap.mux.Unlock()
	if exists {
		return
	}
	// Fetch URL
	body, urls, err := fetcher.Fetch(url)
	// Error check
	if err != nil {
		logger.Println(err)
		return
	}
	// Mark URL as visited
	cmap.mux.Lock()
	cmap.visited[url] = true
	cmap.mux.Unlock()
	// Print display
	logger.Printf("found: %s %q\n", url, body)
	// Recursive calls to new URLs
	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	logger = *log.New(os.Stdout, "", 0)
	cmap.visited = make(map[string]bool)
	wg = sync.WaitGroup{}
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
