package main

import (
	"fmt"
	"time"
)

func main() {
	num := make(chan int)
	str := make(chan string)

	fmt.Println("1st routine")

	go func() {
		time.Sleep(1 * time.Second)
		num <- 1
	}()

	fmt.Println("2nd routine")

	go func() {
		time.Sleep(2 * time.Second)
		str <- "Done"
	}()

	for i := 0; i < 2; i++ {
		select {
		case first := <-num:
			fmt.Println(first)
		case second := <-str:
			fmt.Println(second)
		}
	}

	// <-num
	// <-str
	fmt.Println("Exit...")
}
