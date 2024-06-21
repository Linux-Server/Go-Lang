package main

import (
	"fmt"
)

func main() {

	my_slice := []int{10, 20, 30, 40}

	for _, vals := range my_slice {
		fmt.Println(vals)
	}

	my_map := map[int]int{1: 10, 2: 20, 3: 30}

	for key, _ := range my_map {
		fmt.Println(key)
	}

	for i, val := range "sam" {
		fmt.Println(string(rune(val)), i)
	}

}
