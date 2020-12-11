package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type workerService struct {
	id      int                // for tracking the worker
	request chan int           // for taking in a request
	results chan<- int         // for writing the results
	ctx     context.Context    // for being able to process cancelation
	cancel  context.CancelFunc // for canceling the worker. Calling cancel() writes to the "Done" chan
}

func newWorkerService(ctx context.Context, cancel context.CancelFunc, id int, results chan<- int) *workerService {
	return &workerService{ctx: ctx, cancel: cancel, id: id, request: make(chan int), results: results}
}

// meant to simulate slow and/or seemingly random services
func (w *workerService) doWork() {
	min := 50                         // sleep for min of 50ms
	max := 300                        // sleep for max of 300ms
	sleep := rand.Intn(max-min) + min // calc the sleep time

	// block until we get a request
	<-w.request
	// block for a certain duration (to simulate latency) or until cancel is called.
	select {
	case <-time.After(time.Millisecond * time.Duration(sleep)):
		// echo back the id of the worker who completed the request
		w.results <- w.id
		return
	case <-w.ctx.Done():
		fmt.Printf("worker: %d canceled\n", w.id)
		return
	}
}

// simulate making request to multiple instances of the same service and canceling the request when one returns
func main() {
	// Setup results chan
	results := make(chan int)
	// Set up n workers
	n := 5
	// Track the workers
	workers := []*workerService{}
	for i := 0; i < n; i++ {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		w := newWorkerService(ctx, cancel, i, results)
		workers = append(workers, w)
		go w.doWork()
	}
	start := time.Now() // Track how long the request takes

	// propigate a request to all the workers
	for _, w := range workers {
		w.request <- 1
	}

	// get result
	workerID := <-results

	// cancel workers
	for _, w := range workers {
		w.cancel()
	}

	elapsed := time.Since(start)
	log.Printf("Request took %s and was completed by worker with ID: %d", elapsed, workerID)
}
