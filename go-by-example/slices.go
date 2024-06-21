// package main

// import "fmt"

// func main() {

// 	// var slices []int
// 	// _ = slices

// 	slice_two := []int{1, 2, 3, 4, 5, 6, 7} // [start,end]

// 	fmt.Println(slice_two)
// 	fmt.Println(slice_two[4:])

// 	slice_two = slice_two[:6]
// 	fmt.Println(slice_two)

// 	slice_two = append(slice_two, 7)
// 	fmt.Println(slice_two)

// 	slice_two[3] = 400
// 	fmt.Println(slice_two)

// 	slice_two = append(slice_two[:3], slice_two[4:]...)
// 	fmt.Println(slice_two)

// 	one := make([]int, 3)
// 	// one[0] = 10
// 	// one[1] = 20
// 	// one[3] = 30
// 	two := []int{10, 20, 30}

// 	//two := one
// 	// two[2] = 100

// 	copy(one, two)

// 	fmt.Println("The one : ", one)
// 	fmt.Println("The two : ", two)

// 	// fmt.Println("The slice is : ", slice_two, slice_two == nil, len(slice_two) == 0)

// 	// slice_two = append(slice_two, 10, 20, 30)

// 	// fmt.Println(slice_two)

// 	// // 	// array_two = append(array_two[:4], array_two[5:]...)

// 	// slice_three := make([]int, 3, 6)
// 	// fmt.Println("original: ", slice_three)
// 	// slice_three = append(slice_three, 100, 300, 400, 500)
// 	// fmt.Println(slice_three)
// 	// fmt.Println("The len is :", len(slice_three), "and capacity is :", cap(slice_three))

// }
