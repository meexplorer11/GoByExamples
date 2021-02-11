package main

import "fmt"

func main() {
	age := 20

	if age <= 18 {
		fmt.Println("Not Working")
	} else if age > 18 && age <= 60 {
		fmt.Println("Working")
	} else {
		fmt.Println("Retired")
	}
}
