package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someExpensiveTask(done chan<- struct{}) {
	fmt.Println("Starting expensive task")
	// Sleep for a random amount of time, up to 5 seconds
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	fmt.Println("Expensive task finished")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go someExpensiveTask(done)
	// do other stuff
	<-done
	fmt.Println("Done")
}
