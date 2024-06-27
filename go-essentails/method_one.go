package main

import "fmt"

func method_one() {
	fmt.Println("hello meth")
	stud_one := Student{name: "Sam", id: 12, marks: Marks{10, 20}}
	fmt.Println(stud_one)
	stud_one.new("ram")
	fmt.Println("old value", stud_one)

}

type Student struct {
	name  string
	id    int
	marks Marks
}

func (val *Student) new(name string) {
	val.name = name
	fmt.Println("the new name : ", val)
}

type Marks struct {
	english int
	maths   int
}
