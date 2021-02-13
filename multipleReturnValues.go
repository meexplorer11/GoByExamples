package main

import "fmt"

func main() {
	a, b, c := iReturnMultipleValues()

	fmt.Println(a, b, c)

	a, _, c = iReturnMultipleValues()

	fmt.Println(a, c)

	varArgsMethod("var args", 5, 9)
}

func iReturnMultipleValues() (int, int, int) {
	return 1, 2, 3
}

func varArgsMethod(a string, nums ...int) {
	fmt.Println(a, "==>", nums)
}
