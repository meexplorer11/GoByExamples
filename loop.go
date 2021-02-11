package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("=============")

	for {
		fmt.Println("Inside an Infinite loop..")
		break
	}
	fmt.Println("=============")

	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(fmt.Sprint(n, " is an odd number"))
	}
}
