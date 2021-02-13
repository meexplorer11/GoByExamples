package main

import "fmt"

func main() {
	a := 1

	fmt.Println(a)

	value(&a)

	fmt.Println(a)
}

func value(iptr *int) {
	*iptr = 0
}
