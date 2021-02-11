package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	fmt.Println(len(a))

	b := [2]int{0, 1}
	fmt.Println(b)

	a = [5]int{3, 4}
	fmt.Println(a)

	for index, element := range a {
		fmt.Println(index, element)
	}
}
