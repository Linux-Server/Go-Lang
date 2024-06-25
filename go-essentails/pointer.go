package main

import "fmt"

func pointers() {
	i := 10

	// fmt.Println("The val is : ", i)
	// fmt.Println("The add is : ", &i)

	derefernce_check_copy(i)
	fmt.Println("result : ", i)

}

func derefernce_check(val *int) {
	fmt.Println("Deref : ", val)
	*val = 11
}

func derefernce_check_copy(val int) {
	fmt.Println("Deref : ", val)
	val = 11
}

// &  ==> reference  ==> memory add
// *   ==> derefernce ==> actual value

//  a:= 10
// &a  // 0x0032
// *&a // 10
