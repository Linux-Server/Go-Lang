package main

import "fmt"

func Interfaces() {
	fmt.Println("Hello Interface")

	bike := Bike{"KTM", "2021"}
	car := Car{1000.12}

	// fmt.Println(bike.getData())
	// fmt.Println(car.getData())
	measure(bike)
	measure(car)

}

func measure(g my_inteface) {
	// fmt.Println(g)
	fmt.Println(g.getData())
}

type my_inteface interface {
	getData() string
}

type Bike struct {
	brand string
	model string
}

type Car struct {
	price float64
}

func (val Bike) getData() string {
	return fmt.Sprintf("This is a Bike with %s and model is %s", val.brand, val.model)
}

func (val Car) getData() string {
	return fmt.Sprintf("This is a Car at price of %f", val.price)
}
