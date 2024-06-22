package main

import "fmt"

func varaidic() {
	fmt.Println("Its varaiadic")
	mySlice := []int{1, 2, 3, 4, 5}
	tester(mySlice...)
}

func tester(data ...int) {

	fmt.Println(data)

	for i, val := range data {
		fmt.Println(i, val)
	}

}
