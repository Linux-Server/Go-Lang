package main

import "fmt"

func methods_check() {
	fmt.Println("hello check")
	val := 10 //main  ===intial =10  ===>

	testers(val)
	fmt.Println("The val is :", val)

	testers_pointer(&val)
	fmt.Println("The val is :", val)

}

func testers(a int) {
	a = 20
	// fmt.Println("Tester val is ", a)
}

func testers_pointer(a *int) {
	*a = 20
	// fmt.Println("Tester point val is ", a)
}
