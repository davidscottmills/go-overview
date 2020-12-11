package main

import "fmt"

func printInput(c <-chan string) {
	// range over c until it is closed
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	c := make(chan string)
	defer close(c)
	go printInput(c)
	c <- "Hello"
	c <- "World"
}
