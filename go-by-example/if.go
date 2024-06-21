package main

import "fmt"

func main() {

	if status := 9; status > 10 {
		fmt.Println("The value is bigger")
	} else if status < 10 {
		fmt.Println("this is smalll")
	} else {
		fmt.Println("This is equal")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
