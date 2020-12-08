package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type workerService struct {
	id      int
	request chan int
	results chan<- int
	ctx     context.Context
	cancel  context.CancelFunc
}

func newWorkerService(ctx context.Context, cancel context.CancelFunc, id int, results chan<- int) *workerService {
	return &workerService{ctx: ctx, cancel: cancel, id: id, request: make(chan int), results: results}
}

func (w *workerService) doWork() {
	min := 50
	max := 300
	sleep := rand.Intn(max-min) + min

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

func main() {
	// Setup results chan
	results := make(chan int)
	// Set up n workers
	n := 5
	workers := []*workerService{}
	for i := 0; i < n; i++ {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		w := newWorkerService(ctx, cancel, i, results)
		workers = append(workers, w)
		go w.doWork()
	}
	start := time.Now()

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
