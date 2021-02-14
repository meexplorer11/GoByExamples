package main

import "fmt"

func main() {
	stream := make(chan string, 3)

	stream <- "first"
	stream <- "second"
	stream <- "third"

	close(stream)

	for item := range stream {
		fmt.Println(item)
	}
}
