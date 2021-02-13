package main

import "fmt"

func main() {
	m := make(map[string]string)

	fmt.Println(m)

	fmt.Println(m["k1"])

	if m["k1"] == "" {
		fmt.Println("Key is not present..")
	}

	a := map[string]string{"k1": "v1", "k2": "v2"}

	val, prs := a["k1"]
	fmt.Println(val, prs)

	val, prs = a["k3"]
	fmt.Println(val, prs)

	for key, value := range a {
		fmt.Println(key, value)
	}
}
