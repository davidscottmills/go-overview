package main

import "fmt"

type worker struct {
	id      int
	jobs    <-chan int
	results chan<- string
}

// We pass the same jobs chan when creating the worker, as we want all workers reading from the same channel
func newWorker(id int, jobs <-chan int, results chan<- string) *worker {
	return &worker{id, jobs, results}
}

func (w *worker) doWork() {
	// read from jobs until jobs is closed
	for job := range w.jobs {
		// Print the worker and job id
		w.results <- fmt.Sprintf("worker id: %d, job: %d", w.id, job)
	}
}

func main() {
	numJobs := 100

	jobs := make(chan int)
	results := make(chan string)

	numWorkers := 3
	// Create and start the workers
	for i := 0; i < numWorkers; i++ {
		w := newWorker(i, jobs, results)
		go w.doWork()
	}

	// Fill up the jobs chan
	go func() {
		for i := 0; i < numJobs; i++ {
			jobs <- i
		}
	}()

	// Read the results
	for i := 0; i < numJobs; i++ {
		// read result
		res := <-results
		fmt.Println(res)
	}

	close(jobs)
}
