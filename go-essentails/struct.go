package main

import (
	"fmt"
)

func structs() {
	fmt.Println("hello struct")
	person_one := Person{name: "sam", age: 20}
	person_two := Person{name: "ray"}
	person_two.age = 55

	fmt.Println(person_one, person_two)

	fmt.Println(*call_person("alice", 33))
}

type Person struct {
	name string
	age  int
}

func call_person(name string, age int) *Person {
	return &Person{name, age}
}
