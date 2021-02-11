package main

import "fmt"

func main() {
	//doesn't declare and initialize a constant in this way
	// const a := 20

	// fmt.Println(a)

	const s string = "Hello"

	//can't reassign a constant
	//s = "change"

	fmt.Println(s)
}
