package main

import (
	"fmt"
	"time"
)

func main() {
	times := time.Now().Weekday()
	switch times {
	case time.Monday, time.Thursday:
		fmt.Println("its true")
	case time.Friday:
		fmt.Println("Its false")
	default:
		fmt.Println("Its a default")
	}

	whatAmI := func(a interface{}) {
		switch b := a.(type) {
		case int:
			fmt.Println("Its a integer")
		case bool:
			fmt.Println("Its a bool")
		default:
			fmt.Println("Its a default", b)

		}
	}

	whatAmI(10)
}
