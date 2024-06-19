package main

import "fmt"

func main() {
	for {
		fmt.Println("hello Infinity")
		break
	}

	i := 20
	for i > 10 {
		fmt.Println("loops execting..")
		i -= 1
	}

	for i := range 10 {
		fmt.Println("the range is : ", i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("The value", j)
	}
}
