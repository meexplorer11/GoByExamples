package main

import "fmt"

func main() {
	//empty string
	var s string
	fmt.Println(s)

	var a int //declared
	a = 10    //initialized
	fmt.Println(a)

	var isValid bool
	isValid = false || true
	fmt.Println(isValid)

	var precision float64
	precision = 2.13
	fmt.Println(precision)

	//more than one variables
	var first, second int = 1, 2
	fmt.Println(first, second)

}
