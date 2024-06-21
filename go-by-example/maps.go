package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := make(map[int]string)
	m1[1] = "sam"
	m1[2] = "ram"
	_, prs := m1[3]
	fmt.Println(prs)

	m2 := map[int]string{1: "true", 2: "false"}
	m2[3] = "true"

	fmt.Println("The map: ", m2)

	delete(m2, 2)

	// clear(m2)

	fmt.Println("The map: ", m2)

	if maps.Equal(m1, m2) {
		fmt.Println("true")
	}

}
