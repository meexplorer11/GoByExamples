package main

import (
	"fmt"
	"reflect"
)

func main() {
	age := 20
	isAdult := false
	//works with switch/case only
	//fmt.Println(age.(type))

	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(isAdult))

	fmt.Println(reflect.ValueOf(age).Kind())
}
