package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Size       int64
	Duration   time.Duration
	Error      error
}

// fetchURL fetches a URL and returns the result
func fetchURL(url string) Result {
	start := time.Now()
	result := Result{URL: url}

	resp, err := http.Get(url)
	if err != nil {
		result.Error = err
		return result
	}
	defer resp.Body.Close()

	// Read the response body
	size, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		result.Error = err
		return result
	}

	result.StatusCode = resp.StatusCode
	result.Size = size
	result.Duration = time.Since(start)
	return result
}

// validateURL checks if a URL is valid
func validateURL(rawURL string) (string, error) {
	// Add scheme if missing
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	// Parse the URL to validate it
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsedURL.String(), nil
}

func main() {

	// List of URLs to fetch
	rawURLs := []string{
		"golang.org",
		"github.com",
		"stackoverflow.com",
		"invalid-url-example",
		"example.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(rawURLs))

	// Start go routine for each raw url
	for _, rawURL := range rawURLs {
		wg.Add(1)
		go func(rawURL string) {
			defer wg.Done()

			// Validate URL
			validURL, err := validateURL(rawURL)
			if err != nil {
				results <- Result{URL: rawURL, Error: err}
				return
			}

			// Fetch the URL
			result := fetchURL(validURL)
			results <- result
		}(rawURL)
	}

	// Start a goroutine to close results channel once all fetches are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results as they come in
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error fetching %s: %v\n", result.URL, result.Error)
			continue
		}

		fmt.Printf("Successfully fetched %s - Status: %d, Size: %d bytes, Time: %v\n",
			result.URL, result.StatusCode, result.Size, result.Duration)
	}

	fmt.Println("All URLs processed")
}
