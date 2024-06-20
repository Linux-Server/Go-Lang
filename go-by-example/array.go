package main

import "fmt"

func main() {

	// Array

	// var testArr [5]int
	// _ = testArr
	// // create an array
	// myArr := [...]int{1, 2, 3}
	// _ = myArr

	// fmt.Println("The array is : ", myArr)

	// a := [4]int{1, 2, 3, 4}
	// fmt.Println(a)
	// a = [...]int{100, 2: 99, 400}
	// fmt.Println(a)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	// twoD = [2][3]int{{0, 1, 20}, {1, 2, 30}}
	// fmt.Println("2d: ", twoD)

	// Slices

	var mySlice []int
	_ = mySlice

	slices := []int{1, 2, 3}
	_ = slices

	slic := make([]int, 10, 14)
	slic = append(slic, 12, 33, 34, 35, 25, 77, 78)
	slic = slic[:5]
	_ = slic

	fmt.Println("The slice is : ", mySlice)
	fmt.Println("The slice is : ", slices)

	fmt.Println(slic)

	sliced := []int{1, 2, 3, 4, 5}
    sliced= sliced[1:]...
	fmt.Println("The slice is : ", sliced)
	// array_two = append(array_two[:4], array_two[5:]...)

}
