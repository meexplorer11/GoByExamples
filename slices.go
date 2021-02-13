package main

import "fmt"

func main() {
	s := make([]int, 3)

	fmt.Println(s)

	s = append(s, 2, 3)

	fmt.Println(s)

	a := s[:5]

	fmt.Println(a)

	b := []string{"h", "e", "l", "l", "o"}

	fmt.Println(b)
}
