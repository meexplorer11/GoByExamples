package main

import "fmt"

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
		fmt.Println("Something else")
	}

	switch {
	case i%2 == 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}
