package main

import "fmt"

func methods() {

	fmt.Println("Methods")
	// stud_one_mark := Marks{english: 97, maths: 88}
	// stud_one := Student{name: "sam", id: 10, marks: stud_one_mark}

	// fmt.Println(stud_one.marks.maths)

	rect_one := rect{10, 5}
	new_rect := rect_one.area()
	new_rect_pointer := rect_one.area_new()

	fmt.Println("Old", rect_one)
	fmt.Println("New", new_rect_pointer)

	fmt.Println("Old", rect_one)
	fmt.Println("New", new_rect)
}

type rect struct {
	width  int
	height int
}

func (val rect) area() rect {
	return rect{val.width + 10, val.height + 10}
}

func (val *rect) area_new() rect {
	return rect{val.width + 10, val.height + 10}
}

// type Student struct {
// 	name  string
// 	id    int
// 	marks Marks
// }

// type Marks struct {
// 	english int
// 	maths   int
// }

// func (val Student) new(name string, id int, mark Marks){

// }
