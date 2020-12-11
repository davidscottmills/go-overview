package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type resultsList struct {
	values []string
	mu     sync.Mutex
}

func getURL(wg *sync.WaitGroup, results *resultsList, url string) {

	min := 50                         // sleep for min of 50ms
	max := 300                        // sleep for max of 300ms
	sleep := rand.Intn(max-min) + min // calc the sleep time

	time.Sleep(time.Duration(sleep))

	// We could use channels and some sort of listener to achieve concurrency safe writes, but a mutex will be fine.
	results.mu.Lock()
	results.values = append(results.values, url)
	results.mu.Unlock()
	// Decrement the counter when the goroutine completes.
	wg.Done()

}

// A more understand Task.WhenAll()
func main() {
	resultsList := &resultsList{values: []string{}, mu: sync.Mutex{}}

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.bing.com/",
		"http://www.fakeblock.com/",
		"http://www.facebook.com/",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go getURL(&wg, resultsList, url)
	}
	// Wait for all workers to complete.
	wg.Wait()

	// Print the results
	resultsList.mu.Lock()
	fmt.Println(resultsList.values)
	resultsList.mu.Unlock()

}
