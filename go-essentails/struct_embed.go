package main

import (
	"fmt"
)

func struct_embed() {
	fmt.Println("Struct embed")

	stud_one := Stud{name: "sam", age: 22, Mark: Mark{english: 88, maths: 78}}
	fmt.Println("The stud : ", stud_one.maths)
	stud_one.getMarks()

}

type Stud struct {
	name string
	age  int
	Mark
}

type Mark struct {
	english float64
	maths   float64
}

func (val Mark) getMarks() {
	fmt.Printf("The marks of eng is %f and  maths is %f: ", val.english, val.maths)
}
