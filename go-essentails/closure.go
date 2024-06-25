package main

import "fmt"

func intSeq() func() int {
	i := 0 //1
	fmt.Println("value of i : ", i)
	return func() int {
		i++
		return i
	}
}

func mains() {

	nextInt := intSeq()
	nextInt1 := intSeq()

	fmt.Println("Current value of next : ", nextInt1()) //1
	fmt.Println("Current value of next : ", nextInt())  // 2

	anonymous_func := func() {
		fmt.Println("Im anonymous")
	}

	anonymous_func()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// newInts := intSeq()
	// fmt.Println(newInts())
}
