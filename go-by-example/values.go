package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true)   // true
	fmt.Println(false && true)  // false
	fmt.Println(true && false)  // false
	fmt.Println(false && false) // false

	fmt.Println(true || true)              // true
	fmt.Println(false || true)             // true
	fmt.Println(true || false)             // true
	fmt.Println(false || false)            // false
	fmt.Println("Exclamation: ", !!!false) // true
	variables()

}

func variables() {
	var names = "ray"
	name := "Sam" // short assignment , not global

	var a, b, c int = 1, 2, 3

	status := "100"
	status = "200"
	one := int(status)

	fmt.Println(name, names, a, b, c, status, one)

}
