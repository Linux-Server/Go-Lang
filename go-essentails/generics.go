package main

import "fmt"

func mygeneric() {
	fmt.Println("What is generics")
	my_map_int := map[string]int{"sam": 10, "ram": 1}
	my_map_float := map[string]float64{"sam": 10.0, "ram": 1.5}

	fmt.Println(Sumofgenerics(my_map_int))
	fmt.Println(Sumofgenerics(my_map_float))

}

// map{sam: 10, ram:1}

func SumofInts(val map[string]int) int {
	var sum int = 0
	for _, v := range val {
		sum += v
	}

	return sum

}

func SumofFloats(val map[string]float64) float64 {
	var sum float64 = 0
	for _, v := range val {
		sum += v
	}

	return sum

}

type my_map_types interface{ int | float64 }

func Sumofgenerics[val my_map_types](data map[string]val) val {
	var sum val = 0
	for _, v := range data {
		sum += v
	}

	return sum
}

// Generic is used for function /variable
// fReplace with generics
